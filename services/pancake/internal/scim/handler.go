package scim

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

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	filter := bson.M{}

	if raw := r.URL.Query().Get("filter"); raw != "" {
		parsed, err := ParseFilter(raw)
		if err != nil {
			fun.BadRequest(err.Error()).Send(w)
			return
		}
		filter = parsed
	}

	users, err := h.repo.FindAll(r.Context(), filter)
	if fun.Bail(w, err) {
		return
	}

	fun.Respond(w, models.SCIMListResponse{
		Schemas:      []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"},
		TotalResults: len(users),
		StartIndex:   1,
		ItemsPerPage: len(users),
		Resources:    users,
	})
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := fun.From(r)

	var user models.SCIMUser
	if fun.BailInto(w, req, &user) {
		return
	}

	if err := h.repo.Create(r.Context(), &user); fun.Bail(w, err) {
		return
	}

	fun.Respond(w, user, http.StatusCreated)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	req := fun.From(r)

	oid, err := primitive.ObjectIDFromHex(req.Path("id").String())
	if err != nil {
		fun.BadRequest("invalid id").Send(w)
		return
	}

	user, err := h.repo.FindByID(r.Context(), oid)
	if fun.Bail(w, err) {
		return
	}

	fun.Respond(w, user)
}

func (h *Handler) ReplaceUser(w http.ResponseWriter, r *http.Request) {
	req := fun.From(r)

	oid, err := primitive.ObjectIDFromHex(req.Path("id").String())
	if err != nil {
		fun.BadRequest("invalid id").Send(w)
		return
	}

	var user models.SCIMUser
	if fun.BailInto(w, req, &user) {
		return
	}

	if err := h.repo.Replace(r.Context(), oid, &user); fun.Bail(w, err) {
		return
	}

	fun.Respond(w, user)
}

func (h *Handler) PatchUser(w http.ResponseWriter, r *http.Request) {
	req := fun.From(r)

	oid, err := primitive.ObjectIDFromHex(req.Path("id").String())
	if err != nil {
		fun.BadRequest("invalid id").Send(w)
		return
	}

	var patchOp models.SCIMPatchOp
	if fun.BailInto(w, req, &patchOp) {
		return
	}

	update := bson.M{}
	for _, op := range patchOp.Operations {
		switch op.Op {
		case "replace", "add":
			if op.Path != "" {
				update[op.Path] = op.Value
			}
		case "remove":
			update[op.Path] = nil
		}
	}

	if len(update) == 0 {
		fun.BadRequest("no valid operations").Send(w)
		return
	}

	if err := h.repo.Update(r.Context(), oid, update); fun.Bail(w, err) {
		return
	}

	fun.NoContent().Send(w)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	req := fun.From(r)

	oid, err := primitive.ObjectIDFromHex(req.Path("id").String())
	if err != nil {
		fun.BadRequest("invalid id").Send(w)
		return
	}

	if err := h.repo.Delete(r.Context(), oid); fun.Bail(w, err) {
		return
	}

	fun.NoContent().Send(w)
}
