package notes

import (
	"fmt"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

type PatchOp struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value any    `json:"value"`
}

var allowedPaths = map[string]string{
	"/title": "title",
	"/color": "color",
	"/emoji": "emoji",
	"/tags":  "tags",
}

func applyPatch(ops []PatchOp) (bson.M, error) {
	set := bson.M{}
	push := bson.M{}
	pull := bson.M{}
	unset := bson.M{}

	for _, op := range ops {
		op.Op = strings.ToLower(op.Op)

		// tags array ops: /tags/- (append), /tags/N (index remove/replace)
		if strings.HasPrefix(op.Path, "/tags/") {
			idx := strings.TrimPrefix(op.Path, "/tags/")

			if idx == "-" {
				if op.Op != "add" {
					return nil, fmt.Errorf("op %q not valid for /tags/-", op.Op)
				}
				push["tags"] = op.Value
				continue
			}

			n, err := strconv.Atoi(idx)
			if err != nil {
				return nil, fmt.Errorf("invalid tags index: %q", idx)
			}
			_ = n

			switch op.Op {
			case "remove":
				set[fmt.Sprintf("tags.%s", idx)] = nil
				pull["tags"] = nil
			case "replace":
				set[fmt.Sprintf("tags.%s", idx)] = op.Value
			default:
				return nil, fmt.Errorf("op %q not supported for array index", op.Op)
			}
			continue
		}

		field, ok := allowedPaths[op.Path]
		if !ok {
			return nil, fmt.Errorf("path %q is not allowed", op.Path)
		}

		switch op.Op {
		case "replace", "add":
			set[field] = op.Value
		case "remove":
			unset[field] = ""
		default:
			return nil, fmt.Errorf("unsupported op: %q", op.Op)
		}
	}

	update := bson.M{}
	if len(set) > 0 {
		update["$set"] = set
	}
	if len(unset) > 0 {
		update["$unset"] = unset
	}
	if len(push) > 0 {
		update["$push"] = push
	}
	if len(pull) > 0 {
		update["$pull"] = pull
	}

	if len(update) == 0 {
		return nil, fmt.Errorf("no valid operations")
	}

	return update, nil
}
