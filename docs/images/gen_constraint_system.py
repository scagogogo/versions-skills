#!/usr/bin/env python3
"""
Generate a flat-style constraint system hierarchy diagram — refined.
Color palette: blue-gray, no purple, square corners, generous spacing.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch

fig, ax = plt.subplots(figsize=(20, 13))
ax.set_xlim(0, 20)
ax.set_ylim(0, 13)
ax.axis("off")
fig.patch.set_facecolor("white")

# Title
ax.text(10, 12.3, "Constraint System", ha="center", va="center", fontsize=22, fontweight="bold", color="#0f172a")
ax.text(10, 11.7, "Grammar hierarchy: Union → Set (AND) → Single → Operator + Operand", ha="center", va="center", fontsize=11, color="#64748b")

def draw_node(x, y, w, h, label, color, font_size=10, bold=True):
    box = FancyBboxPatch((x - w / 2, y - h / 2), w, h,
        boxstyle="square,pad=0.03", facecolor=color, edgecolor=color, lw=1.2, alpha=0.92, zorder=5)
    ax.add_patch(box)
    weight = "bold" if bold else "normal"
    ax.text(x, y, label, ha="center", va="center", fontsize=font_size, fontweight=weight, color="white", zorder=6, family="monospace")

def draw_leaf(x, y, w, h, label, color):
    box = FancyBboxPatch((x - w / 2, y - h / 2), w, h,
        boxstyle="square,pad=0.02", facecolor="white", edgecolor=color, lw=1, zorder=5)
    ax.add_patch(box)
    ax.text(x, y, label, ha="center", va="center", fontsize=9, color="#334155", zorder=6, family="monospace")

def draw_arrow(x1, y1, x2, y2, color="#94a3b8"):
    ax.annotate("", xy=(x2, y2), xytext=(x1, y1),
        arrowprops=dict(arrowstyle="->", color=color, lw=1.5), zorder=3)

# ═══ Left side: Grammar tree ═══
# Top
draw_node(5, 10.3, 3.5, 0.6, "ConstraintExpr", "#0f172a", font_size=12)

# Level 2
draw_node(2.5, 8.5, 3.0, 0.55, "ConstraintUnion", "#2563eb", font_size=10)
draw_arrow(4.0, 10.0, 2.5, 8.78)

draw_node(5.0, 8.5, 3.0, 0.55, "Constraint (single)", "#16a34a", font_size=10)
draw_arrow(5.0, 10.0, 5.0, 8.78)

draw_node(7.5, 8.5, 3.0, 0.55, "NegatedConstraint", "#dc2626", font_size=10)
draw_arrow(6.0, 10.0, 7.5, 8.78)

# ConstraintUnion → ConstraintSet
draw_node(2.5, 7.0, 3.0, 0.55, "ConstraintSet (AND)", "#0ea5e9", font_size=9.5)
draw_arrow(2.5, 8.22, 2.5, 7.28)

# ConstraintSet → Constraint
draw_leaf(2.5, 5.9, 2.5, 0.40, "Constraint (single)", "#0ea5e9")
draw_arrow(2.5, 6.72, 2.5, 6.10)

# Constraint → Operator + VersionOperand
draw_node(4.0, 7.0, 2.2, 0.55, "Operator", "#0891b2", font_size=10)
draw_arrow(5.5, 8.22, 4.3, 7.28)

draw_node(6.0, 7.0, 2.6, 0.55, "VersionOperand", "#ea580c", font_size=10)
draw_arrow(5.5, 8.22, 5.7, 7.28)

# Operator leaves
ops = ["=", "!=", ">", ">=", "<", "<="]
leaf_w = 0.65
leaf_h = 0.38
leaf_gap = 0.18
total_w = len(ops) * leaf_w + (len(ops) - 1) * leaf_gap
leaf_start_x = 4.0 - total_w / 2

for i, op in enumerate(ops):
    lx = leaf_start_x + i * (leaf_w + leaf_gap) + leaf_w / 2
    ly = 5.7
    draw_leaf(lx, ly, leaf_w, leaf_h, op, "#0891b2")
    draw_arrow(4.0, 6.72, lx, ly + leaf_h / 2 + 0.05)

# VersionOperand leaves
vo_leaves = ["e.g. 1.0.0", "e.g. 2.3.0-beta"]
for i, vl in enumerate(vo_leaves):
    ly = 5.9 - i * 0.60
    draw_leaf(6.0, ly, 2.2, 0.40, vl, "#ea580c")
    draw_arrow(6.0, 6.72, 6.0, ly + 0.20 + 0.05)

# NegatedConstraint → inner ConstraintExpr
draw_node(7.5, 7.0, 2.8, 0.55, "inner ConstraintExpr", "#64748b", font_size=9, bold=False)
draw_arrow(7.5, 8.22, 7.5, 7.28)

# Self-reference
ax.annotate("", xy=(9.0, 10.3), xytext=(9.0, 7.0),
    arrowprops=dict(arrowstyle="->", color="#64748b", lw=1.2, linestyle="dashed"), zorder=3)
ax.text(9.4, 8.6, "recursive", ha="center", va="center", fontsize=8, color="#64748b", style="italic", rotation=90)

# ═══ Right side: Operators + Examples ═══
# Operators section
ax.text(14.5, 10.5, "Operators", ha="center", va="center", fontsize=13, fontweight="bold", color="#0f172a")

operators_full = ["=", "!=", ">", ">=", "<", "<=", "^", "~", "1.x"]
cols = 3
for i, op in enumerate(operators_full):
    row = i // cols
    col = i % cols
    x = 13.0 + col * 1.8
    y = 9.7 - row * 0.75
    draw_leaf(x, y, 0.85, 0.42, op, "#2563eb")

# Examples section
ax.text(14.5, 7.6, "Examples", ha="center", va="center", fontsize=13, fontweight="bold", color="#0f172a")

examples = [
    (">=1.0.0", "At least 1.0.0"),
    ("^1.2.3", "Compatible with 1.2.3"),
    ("~1.2", "Approximately 1.2"),
    ("1.x", "Any 1.* version"),
    (">=1.0.0,<2.0.0", "Range: 1.x"),
    (">=1.0.0 || >=3.0.0", "1.x or 3.x+"),
    ("!=1.2.3", "Not 1.2.3"),
    ("<2.0.0", "Below 2.0.0"),
]

for i, (expr, desc) in enumerate(examples):
    y = 6.9 - i * 0.65
    # Expression in a tag
    expr_box = FancyBboxPatch((11.5, y - 0.18), 2.6, 0.36,
        boxstyle="square,pad=0.02", facecolor="#2563eb", edgecolor="#2563eb", lw=0, alpha=0.10, zorder=4)
    ax.add_patch(expr_box)
    ax.text(12.8, y, expr, ha="center", va="center", fontsize=9.5, color="#1e40af", fontweight="bold", family="monospace", zorder=5)
    # Description
    ax.text(14.5, y, desc, ha="left", va="center", fontsize=9, color="#475569", zorder=5)

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/constraint-system.png", dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("constraint-system.png saved")
