package verfmt

import (
	"github.com/scagogogo/versions-skills"
)

// FormatConstraint 将 Constraint 转换为 map
func FormatConstraint(c *versions.Constraint) map[string]interface{} {
	if c == nil {
		return nil
	}
	result := map[string]interface{}{
		"operator": string(c.Operator),
		"string":   c.String(),
	}
	if c.Version != nil {
		result["version"] = c.Version.RawString()
	}
	return result
}

// FormatConstraintSet 将 ConstraintSet 转换为 map
func FormatConstraintSet(cs *versions.ConstraintSet) map[string]interface{} {
	if cs == nil {
		return nil
	}
	constraints := make([]map[string]interface{}, len(cs.Constraints))
	for i, c := range cs.Constraints {
		constraints[i] = FormatConstraint(&c)
	}
	return map[string]interface{}{
		"string":      cs.String(),
		"len":         cs.Len(),
		"constraints": constraints,
	}
}

// FormatConstraintUnion 将 ConstraintUnion 转换为 map
func FormatConstraintUnion(cu *versions.ConstraintUnion) map[string]interface{} {
	if cu == nil {
		return nil
	}
	sets := make([]map[string]interface{}, len(cu.Sets))
	for i, cs := range cu.Sets {
		sets[i] = FormatConstraintSet(cs)
	}
	return map[string]interface{}{
		"string": cu.String(),
		"sets":   sets,
	}
}
