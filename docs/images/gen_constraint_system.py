#!/usr/bin/env python3
"""
Generate a refined constraint system hierarchy diagram with modern styling.
Features: rounded corners, gradient fills, shadows, elegant typography.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch
import numpy as np

def draw_shadow(ax, x, y, w, h, radius=0.12, offset=0.04, alpha=0.1):
    shadow = FancyBboxPatch(
        (x + offset, y - offset), w, h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor="#000000", edgecolor="none", alpha=alpha, zorder=3
    )
    ax.add_patch(shadow)

def draw_gradient_box(ax, x, y, w, h, color, radius=0.12, alpha=0.92):
    box = FancyBboxPatch(
        (x, y), w, h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor=color, edgecolor="none", alpha=alpha, zorder=5
    )
    ax.add_patch(box)
    highlight_h = h * 0.35
    highlight = FancyBboxPatch(
        (x, y + h - highlight_h), w, highlight_h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor="#ffffff", edgecolor="none", alpha=0.18, zorder=6
    )
    ax.add_patch(highlight)

def draw_node(ax, x, y, w, h, label, color, font_size=10, bold=True, radius=0.12):
    draw_shadow(ax, x - w / 2, y - h / 2, w, h, radius=radius, offset=0.05, alpha=0.12)
    draw_gradient_box(ax, x - w / 2, y - h / 2, w, h, color, radius=radius)
    weight = "bold" if bold else "medium"
    ax.text(x, y, label, ha="center", va="center", fontsize=font_size, fontweight=weight, color="white", zorder=7, family="monospace")

def draw_leaf(ax, x, y, w, h, label, color, radius=0.1):
    draw_shadow(ax, x - w / 2, y - h / 2, w, h, radius=radius, offset=0.03, alpha=0.08)
    box = FancyBboxPatch(
        (x - w / 2, y - h / 2), w, h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor="white", edgecolor=color, lw=1.5, zorder=5
    )
    ax.add_patch(box)
    ax.text(x, y, label, ha="center", va="center", fontsize=9, color="#334155", zorder=6, family="monospace", weight="medium")

def draw_arrow(ax, x1, y1, x2, y2, color="#94a3b8"):
    ax.annotate("", xy=(x2, y2), xytext=(x1, y1),
        arrowprops=dict(arrowstyle="->,head_width=0.3,head_length=0.4", color=color, lw=1.8), zorder=3)

# ─── Drawing ───
fig, ax = plt.subplots(figsize=(22, 14))
ax.set_xlim(0, 22)
ax.set_ylim(0, 14)
ax.axis("off")
fig.patch.set_facecolor("#fafbfc")

# Background grid
for i in range(0, 23, 2):
    ax.axvline(i, color="#f1f5f9", lw=0.5, zorder=0)
for i in range(0, 15, 2):
    ax.axhline(i, color="#f1f5f9", lw=0.5, zorder=0)

# Title
ax.text(11, 13.3, "Constraint System", ha="center", va="center", fontsize=26, fontweight="bold", color="#0f172a", family="sans-serif")
ax.text(11, 12.6, "Grammar hierarchy: Union → Set (AND) → Single → Operator + Operand", ha="center", va="center", fontsize=12, color="#64748b", family="sans-serif")

# ═══ Left side: Grammar tree ═══
# Top node
draw_node(ax, 5, 11.2, 3.8, 0.7, "ConstraintExpr", "#0f172a", font_size=12, radius=0.15)

# Level 2 nodes
draw_node(ax, 2.5, 9.2, 3.2, 0.6, "ConstraintUnion", "#3b82f6", font_size=10, radius=0.12)
draw_arrow(ax, 4.2, 10.85, 2.5, 9.5)

draw_node(ax, 5.0, 9.2, 3.4, 0.6, "Constraint (single)", "#10b981", font_size=10, radius=0.12)
draw_arrow(ax, 5.0, 10.85, 5.0, 9.5)

draw_node(ax, 7.5, 9.2, 3.4, 0.6, "NegatedConstraint", "#ef4444", font_size=10, radius=0.12)
draw_arrow(ax, 5.8, 10.85, 7.5, 9.5)

# ConstraintUnion → ConstraintSet
draw_node(ax, 2.5, 7.7, 3.4, 0.6, "ConstraintSet (AND)", "#06b6d4", font_size=9.5, radius=0.12)
draw_arrow(ax, 2.5, 8.9, 2.5, 8.0)

# ConstraintSet → Constraint
draw_leaf(ax, 2.5, 6.5, 2.8, 0.45, "Constraint (single)", "#06b6d4")
draw_arrow(ax, 2.5, 7.4, 2.5, 6.73)

# Constraint → Operator + VersionOperand
draw_node(ax, 4.2, 7.7, 2.4, 0.6, "Operator", "#0ea5e9", font_size=10, radius=0.12)
draw_arrow(ax, 5.6, 8.9, 4.5, 8.0)

draw_node(ax, 6.0, 7.7, 2.8, 0.6, "VersionOperand", "#f97316", font_size=10, radius=0.12)
draw_arrow(ax, 5.6, 8.9, 5.8, 8.0)

# Operator leaves
ops = ["=", "!=", ">", ">=", "<", "<="]
leaf_w = 0.75
leaf_h = 0.42
leaf_gap = 0.2
total_w = len(ops) * leaf_w + (len(ops) - 1) * leaf_gap
leaf_start_x = 4.2 - total_w / 2

for i, op in enumerate(ops):
    lx = leaf_start_x + i * (leaf_w + leaf_gap) + leaf_w / 2
    ly = 6.3
    draw_leaf(ax, lx, ly, leaf_w, leaf_h, op, "#0ea5e9", radius=0.08)
    draw_arrow(ax, 4.2, 7.4, lx, ly + leaf_h / 2 + 0.08)

# VersionOperand leaves
vo_leaves = ["e.g. 1.0.0", "e.g. 2.3.0-beta"]
for i, vl in enumerate(vo_leaves):
    ly = 6.5 - i * 0.65
    draw_leaf(ax, 6.0, ly, 2.4, 0.45, vl, "#f97316", radius=0.1)
    draw_arrow(ax, 6.0, 7.4, 6.0, ly + 0.23 + 0.08)

# NegatedConstraint → inner ConstraintExpr
draw_node(ax, 7.5, 7.7, 3.0, 0.6, "inner ConstraintExpr", "#64748b", font_size=9, bold=False, radius=0.12)
draw_arrow(ax, 7.5, 8.9, 7.5, 8.0)

# Self-reference (recursive)
ax.annotate("", xy=(9.5, 11.2), xytext=(9.5, 7.7),
    arrowprops=dict(arrowstyle="->,head_width=0.3,head_length=0.4", color="#64748b", lw=1.5, linestyle="dashed"), zorder=3)
ax.text(9.9, 9.4, "recursive", ha="center", va="center", fontsize=9, color="#64748b", style="italic", rotation=90)

# ═══ Right side: Operators + Examples ═══
# Operators section
ax.text(15.5, 11.0, "Operators", ha="center", va="center", fontsize=14, fontweight="bold", color="#0f172a", family="sans-serif")

operators_full = ["=", "!=", ">", ">=", "<", "<=", "^", "~", "1.x"]
cols = 3
op_radius = 0.12
for i, op in enumerate(operators_full):
    row = i // cols
    col = i % cols
    x = 14.0 + col * 2.0
    y = 10.1 - row * 0.85
    draw_leaf(ax, x, y, 0.95, 0.5, op, "#3b82f6", radius=op_radius)

# Examples section
ax.text(15.5, 7.9, "Examples", ha="center", va="center", fontsize=14, fontweight="bold", color="#0f172a", family="sans-serif")

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
    y = 7.2 - i * 0.7

    # Expression tag — pill style
    expr_w = 2.8
    expr_h = 0.42
    draw_shadow(ax, 12.2, y - expr_h / 2, expr_w, expr_h, radius=0.2, offset=0.03, alpha=0.08)
    expr_box = FancyBboxPatch(
        (12.2, y - expr_h / 2), expr_w, expr_h,
        boxstyle="round,pad=0,rounding_size=0.2",
        facecolor="#3b82f6", edgecolor="none", alpha=0.12, zorder=4
    )
    ax.add_patch(expr_box)
    ax.text(13.6, y, expr, ha="center", va="center", fontsize=10, color="#1e40af", fontweight="bold", family="monospace", zorder=5)

    # Description
    ax.text(15.3, y, desc, ha="left", va="center", fontsize=10, color="#475569", zorder=5, family="sans-serif", weight="medium")

# Decorative corner accents
corner_size = 1.0
corner_color = "#e2e8f0"
ax.plot([0.5, 0.5], [13 - corner_size, 13], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([0.5, 0.5 + corner_size], [13, 13], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([21.5, 21.5], [1 + corner_size, 1], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([21.5, 21.5 - corner_size], [1, 1], color=corner_color, lw=3, solid_capstyle='round')

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/constraint-system.png", dpi=150, bbox_inches="tight", facecolor="#fafbfc")
plt.close()
print("constraint-system.png saved")