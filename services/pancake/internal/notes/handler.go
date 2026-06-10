package notes

import (
	"net/http"
	"pancake/models"

	"github.com/MintzyG/fun"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	req := fun.From(r)
	filter := bson.M{}

	if tag := req.Query("tag").String(); tag != "" {
		filter["tags"] = bson.M{"$in": []string{tag}}
	}

	notes, err := h.repo.FindAll(r.Context(), filter)
	if fun.Bail(w, err) {
		return
	}

	fun.Respond(w, notes)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	req := fun.From(r)

	var note models.Note
	if fun.BailInto(w, req, &note) {
		return
	}

	err := h.repo.Create(r.Context(), &note)
	if fun.Bail(w, err) {
		return
	}

	fun.Respond(w, note, http.StatusCreated)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	req := fun.From(r)

	oid, err := primitive.ObjectIDFromHex(req.Path("id").String())
	if err != nil {
		fun.BadRequest("invalid id").Send(w)
		return
	}

	note, err := h.repo.FindByID(r.Context(), oid)
	if fun.Bail(w, err) {
		return
	}

	fun.Respond(w, note)
}

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	req := fun.From(r)

	oid, err := primitive.ObjectIDFromHex(req.Path("id").String())
	if err != nil {
		fun.BadRequest("invalid id").Send(w)
		return
	}

	var patch map[string]any
	if fun.BailInto(w, req, &patch) {
		return
	}

	allowed := map[string]bool{"title": true, "color": true, "emoji": true, "tags": true}
	update := bson.M{}
	for k, v := range patch {
		if allowed[k] {
			update[k] = v
		}
	}

	if len(update) == 0 {
		fun.BadRequest("no valid fields to update").Send(w)
		return
	}

	err = h.repo.Update(r.Context(), oid, update)
	if fun.Bail(w, err) {
		return
	}

	fun.NoContent().Send(w)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	req := fun.From(r)

	oid, err := primitive.ObjectIDFromHex(req.Path("id").String())
	if err != nil {
		fun.BadRequest("invalid id").Send(w)
		return
	}

	err = h.repo.Delete(r.Context(), oid)
	if fun.Bail(w, err) {
		return
	}

	fun.NoContent().Send(w)
}
