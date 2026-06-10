package scim

import (
	"fmt"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// ParseFilter converts a SCIM filter string (RFC 7644 §3.4.2.2) to a bson.M query.
// Supports single-attribute expressions: attrPath op value
// Operators: eq ne co sw ew gt lt ge le
func ParseFilter(filter string) (bson.M, error) {
	filter = strings.TrimSpace(filter)
	if filter == "" {
		return bson.M{}, nil
	}

	parts := strings.SplitN(filter, " ", 3)
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid filter: %q", filter)
	}

	attr := parts[0]
	op := strings.ToLower(parts[1])
	rawVal := strings.TrimSpace(parts[2])

	val := parseValue(rawVal)

	switch op {
	case "eq":
		return bson.M{attr: val}, nil
	case "ne":
		return bson.M{attr: bson.M{"$ne": val}}, nil
	case "co":
		return bson.M{attr: bson.M{"$regex": fmt.Sprintf("%v", val), "$options": "i"}}, nil
	case "sw":
		return bson.M{attr: bson.M{"$regex": fmt.Sprintf("^%v", val), "$options": "i"}}, nil
	case "ew":
		return bson.M{attr: bson.M{"$regex": fmt.Sprintf("%v$", val), "$options": "i"}}, nil
	case "gt":
		return bson.M{attr: bson.M{"$gt": val}}, nil
	case "lt":
		return bson.M{attr: bson.M{"$lt": val}}, nil
	case "ge":
		return bson.M{attr: bson.M{"$gte": val}}, nil
	case "le":
		return bson.M{attr: bson.M{"$lte": val}}, nil
	default:
		return nil, fmt.Errorf("unsupported operator: %q", op)
	}
}

func parseValue(raw string) any {
	if strings.HasPrefix(raw, `"`) && strings.HasSuffix(raw, `"`) {
		return raw[1 : len(raw)-1]
	}
	if raw == "true" {
		return true
	}
	if raw == "false" {
		return false
	}
	if n, err := strconv.ParseInt(raw, 10, 64); err == nil {
		return n
	}
	if f, err := strconv.ParseFloat(raw, 64); err == nil {
		return f
	}
	return raw
}
