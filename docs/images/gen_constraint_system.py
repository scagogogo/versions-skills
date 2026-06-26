#!/usr/bin/env python3
"""
Generate a constraint expression diagram showing the grammar
and examples of version constraints.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch
import numpy as np

fig, ax = plt.subplots(figsize=(18, 9))
ax.set_xlim(0, 18)
ax.set_ylim(0, 9)
ax.axis("off")
fig.patch.set_facecolor("white")

# Title
ax.text(9, 8.5, "Constraint Expression System", ha="center", va="center",
        fontsize=18, fontweight="bold", color="#1e293b", family="monospace")
ax.text(9, 8.0, "npm-style version constraints with full boolean logic",
        ha="center", va="center", fontsize=10, color="#64748b")

# Grammar hierarchy
grammar_items = [
    {"x": 4.5, "y": 6.5, "w": 7, "h": 1.0, "color": "#8b5cf6", "edge": "#6d28d9",
     "title": "ConstraintUnion (OR)", "detail": ">=1.0.0 || >=3.0.0"},
    {"x": 4.5, "y": 5.0, "w": 7, "h": 1.0, "color": "#3b82f6", "edge": "#1d4ed8",
     "title": "ConstraintSet (AND)", "detail": ">=1.0.0,<2.0.0"},
    {"x": 4.5, "y": 3.5, "w": 7, "h": 1.0, "color": "#10b981", "edge": "#059669",
     "title": "Constraint (Single)", "detail": ">=1.0.0"},
    {"x": 4.5, "y": 2.0, "w": 7, "h": 1.0, "color": "#f59e0b", "edge": "#d97706",
     "title": "Operator + Version", "detail": ">=  +  1.0.0"},
]

for item in grammar_items:
    box = FancyBboxPatch(
        (item["x"] - item["w"]/2, item["y"] - item["h"]/2),
        item["w"], item["h"],
        boxstyle="round,pad=0.1", facecolor=item["color"],
        edgecolor=item["edge"], lw=2, alpha=0.85, zorder=4
    )
    ax.add_patch(box)
    ax.text(item["x"] - 2.5, item["y"] + 0.15, item["title"],
            ha="left", va="center", fontsize=11, fontweight="bold", color="white", zorder=5)
    ax.text(item["x"] + 2.5, item["y"] - 0.15, item["detail"],
            ha="right", va="center", fontsize=9, color="white", alpha=0.85, zorder=5,
            family="monospace")

# Arrows between grammar levels
for i in range(len(grammar_items) - 1):
    ax.annotate("", xy=(4.5, grammar_items[i+1]["y"] + grammar_items[i+1]["h"]/2),
                xytext=(4.5, grammar_items[i]["y"] - grammar_items[i]["h"]/2),
                arrowprops=dict(arrowstyle="-|>", color="#94a3b8", lw=1.5))

# Operators
operators = ["=", "!=", ">", ">=", "<", "<=", "^", "~", "1.x"]
ax.text(13.5, 6.8, "Operators", ha="center", va="center",
        fontsize=12, fontweight="bold", color="#1e293b")
for i, op in enumerate(operators):
    row = i // 3
    col = i % 3
    x = 12.0 + col * 1.5
    y = 6.0 - row * 0.8
    box = FancyBboxPatch(
        (x - 0.55, y - 0.3), 1.1, 0.6,
        boxstyle="round,pad=0.05", facecolor="#f0f9ff",
        edgecolor="#3b82f6", lw=1.5, zorder=4
    )
    ax.add_patch(box)
    ax.text(x, y, op, ha="center", va="center", fontsize=10,
            fontweight="bold", color="#1e40af", zorder=5, family="monospace")

# Examples
examples = [
    (">=1.0.0", "At least 1.0.0"),
    ("^1.2.3", "Compatible with 1.2.3"),
    ("~1.2", "Approximately 1.2"),
    ("1.x", "Any 1.* version"),
    (">=1.0.0,<2.0.0", "Range: 1.x"),
    (">=1.0.0 || >=3.0.0", "1.x or 3.x+"),
]

ax.text(13.5, 3.5, "Examples", ha="center", va="center",
        fontsize=12, fontweight="bold", color="#1e293b")

for i, (expr, desc) in enumerate(examples):
    y = 2.8 - i * 0.55
    ax.text(11.5, y, expr, ha="left", va="center", fontsize=8.5,
            color="#1e40af", family="monospace", fontweight="bold")
    ax.text(14.5, y, desc, ha="left", va="center", fontsize=8, color="#64748b")

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/constraint-system.png",
            dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("constraint-system.png saved")
