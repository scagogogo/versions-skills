#!/usr/bin/env python3
"""
Generate a data flow diagram showing how version strings flow through
the library's processing pipeline.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch
import numpy as np

fig, ax = plt.subplots(figsize=(20, 7))
ax.set_xlim(0, 20)
ax.set_ylim(0, 7)
ax.axis("off")
fig.patch.set_facecolor("white")

stages = [
    {
        "x": 2.0, "y": 3.5, "w": 2.6, "h": 2.0,
        "color": "#3b82f6", "title": "INPUT",
        "items": ["Version String", "File", "CLI Args", "MCP Request"],
    },
    {
        "x": 5.6, "y": 3.5, "w": 2.6, "h": 2.0,
        "color": "#8b5cf6", "title": "PARSE",
        "items": ["NewVersion()", "Coerce()", "MustParse()", "Validate()"],
    },
    {
        "x": 9.2, "y": 3.5, "w": 2.6, "h": 2.0,
        "color": "#10b981", "title": "PROCESS",
        "items": ["Compare / Sort", "Group / Filter", "Constraint Check", "Range Query"],
    },
    {
        "x": 12.8, "y": 3.5, "w": 2.6, "h": 2.0,
        "color": "#f59e0b", "title": "TRANSFORM",
        "items": ["Bump / With*", "Build", "Negate", "Set Operations"],
    },
    {
        "x": 16.4, "y": 3.5, "w": 2.6, "h": 2.0,
        "color": "#ef4444", "title": "OUTPUT",
        "items": ["Structured Result", "Visualization", "File Write", "JSON Response"],
    },
]

for stage in stages:
    box = FancyBboxPatch(
        (stage["x"] - stage["w"]/2, stage["y"] - stage["h"]/2),
        stage["w"], stage["h"],
        boxstyle="round,pad=0.12", facecolor=stage["color"],
        edgecolor="white", lw=2, alpha=0.9, zorder=4
    )
    ax.add_patch(box)

    ax.text(stage["x"], stage["y"] + stage["h"]/2 - 0.3, stage["title"],
            ha="center", va="center", fontsize=11, fontweight="bold",
            color="white", zorder=5, family="monospace")

    for i, item in enumerate(stage["items"]):
        ax.text(stage["x"], stage["y"] + 0.3 - i * 0.38, item,
                ha="center", va="center", fontsize=8.5, color="white",
                alpha=0.9, zorder=5, family="monospace")

# Arrows between stages
for i in range(len(stages) - 1):
    x_from = stages[i]["x"] + stages[i]["w"]/2
    x_to = stages[i+1]["x"] - stages[i+1]["w"]/2
    ax.annotate(
        "", xy=(x_to, stages[i]["y"]), xytext=(x_from, stages[i]["y"]),
        arrowprops=dict(arrowstyle="-|>", color="#94a3b8", lw=2.5),
        zorder=3
    )

# Version object description
ax.text(10, 1.2, "Version{Raw, Prefix, VersionNumbers, Suffix, Metadata, PublicTime}",
        ha="center", va="center", fontsize=9, color="#64748b",
        style="italic", family="monospace",
        bbox=dict(boxstyle="round,pad=0.3", facecolor="#f1f5f9", edgecolor="#cbd5e1"))

ax.annotate("", xy=(5.6, 2.5), xytext=(4.5, 2.5),
            arrowprops=dict(arrowstyle="-|>", color="#8b5cf6", lw=1.5, ls="--"))
ax.annotate("", xy=(14.5, 2.5), xytext=(13.5, 2.5),
            arrowprops=dict(arrowstyle="-|>", color="#f59e0b", lw=1.5, ls="--"))

ax.text(5.0, 2.1, "Version Object", ha="center", va="center", fontsize=7.5,
        color="#8b5cf6", style="italic")
ax.text(14.0, 2.1, "Version Object", ha="center", va="center", fontsize=7.5,
        color="#f59e0b", style="italic")

# Title
ax.text(10, 6.3, "versions-skills  --  Data Flow Pipeline", ha="center", va="center",
        fontsize=18, fontweight="bold", color="#1e293b", family="monospace")

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/data-flow.png",
            dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("data-flow.png saved")
