package verfmt

import (
	"sort"

	"github.com/scagogogo/versions-skills"
)

// FormatVersionGroup 将 VersionGroup 转换为 map
func FormatVersionGroup(g *versions.VersionGroup) map[string]interface{} {
	if g == nil {
		return nil
	}

	latest := g.GetLatest()
	oldest := g.GetOldest()
	latestStable := g.LatestStable()
	latestPrerelease := g.LatestPrerelease()

	result := map[string]interface{}{
		"id":                 g.ID(),
		"count":              g.Count(),
		"versions":           FormatVersionStrings(g.SortVersions()),
	}

	if latest != nil {
		result["latest"] = latest.RawString()
	}
	if oldest != nil {
		result["oldest"] = oldest.RawString()
	}
	if latestStable != nil {
		result["latest_stable"] = latestStable.RawString()
	}
	if latestPrerelease != nil {
		result["latest_prerelease"] = latestPrerelease.RawString()
	}

	return result
}

// FormatVersionGroupMap 将分组 map 转换为有序的 []map[string]interface{}
func FormatVersionGroupMap(groupMap map[string]*versions.VersionGroup) []map[string]interface{} {
	// 按 ID 排序
	ids := make([]string, 0, len(groupMap))
	for id := range groupMap {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	result := make([]map[string]interface{}, 0, len(ids))
	for _, id := range ids {
		result = append(result, FormatVersionGroup(groupMap[id]))
	}

	return result
}
