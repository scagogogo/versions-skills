#!/usr/bin/env python3
"""
Generate a suffix weight ordering diagram showing the semantic
ordering of version suffixes.
"""

import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch
import numpy as np

fig, ax = plt.subplots(figsize=(18, 5))
ax.set_xlim(0, 18)
ax.set_ylim(0, 5)
ax.axis("off")
fig.patch.set_facecolor("white")

# ─── Suffix weight data ───
suffixes = [
    ("dev", 50, "#ef4444"),
    ("snapshot", 60, "#f97316"),
    ("nightly", 70, "#f59e0b"),
    ("alpha", 100, "#eab308"),
    ("beta", 200, "#84cc16"),
    ("milestone", 300, "#22c55e"),
    ("rc", 400, "#10b981"),
    ("final/ga", 500, "#3b82f6"),
    ("sp", 600, "#6366f1"),
    ("patch", 700, "#8b5cf6"),
    ("post", 800, "#a855f7"),
]

# Draw as a horizontal ladder
y_line = 2.5
x_start = 1.5
x_end = 16.5
weight_min = suffixes[0][1]
weight_max = suffixes[-1][1]

# Main axis line
ax.plot([x_start - 0.3, x_end + 0.3], [y_line, y_line], color="#cbd5e1", lw=2, zorder=1)

# Arrow at end
ax.annotate("", xy=(x_end + 0.5, y_line), xytext=(x_end + 0.1, y_line),
            arrowprops=dict(arrowstyle="-|>", color="#94a3b8", lw=2))

for label, weight, color in suffixes:
    # Position on the axis
    frac = (weight - weight_min) / (weight_max - weight_min)
    x = x_start + frac * (x_end - x_start)

    # Dot on axis
    ax.plot(x, y_line, 'o', color=color, markersize=10, zorder=3)

    # Label above
    ax.text(x, y_line + 0.55, label, ha="center", va="center",
            fontsize=10, fontweight="bold", color=color)

    # Weight value below
    ax.text(x, y_line - 0.45, str(weight), ha="center", va="center",
            fontsize=8, color="#94a3b8")

    # Small tick
    ax.plot([x, x], [y_line - 0.15, y_line + 0.15], color=color, lw=2, zorder=2)

# Labels
ax.text(x_start - 0.5, y_line, "Lighter", ha="right", va="center",
        fontsize=9, color="#94a3b8", style="italic")
ax.text(x_end + 0.8, y_line, "Heavier", ha="left", va="center",
        fontsize=9, color="#94a3b8", style="italic")

# Title
ax.text(9, 4.2, "Suffix Weight Ordering — Semantic Comparison Priority", ha="center",
        va="center", fontsize=16, fontweight="bold", color="#1e293b")
ax.text(9, 3.7, "dev < snapshot < nightly < alpha < beta < milestone < rc < final/ga < sp < patch < post",
        ha="center", va="center", fontsize=10, color="#64748b")

# Example
ax.text(9, 1.2, 'Example: 1.0.0-alpha < 1.0.0-beta < 1.0.0-rc < 1.0.0 < 1.0.0-sp1',
        ha="center", va="center", fontsize=10, color="#475569",
        bbox=dict(boxstyle="round,pad=0.3", facecolor="#f0f9ff", edgecolor="#bfdbfe"))

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/suffix-weight.png",
            dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("✅ suffix-weight.png saved")
