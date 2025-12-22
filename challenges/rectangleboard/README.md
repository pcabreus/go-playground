# Layered Rectangles on a Grid (Live Coding)

You are given a 2D board (grid) of size W x H. You can place axis-aligned rectangles on the board. Rectangles may overlap, and the one added most recently should be visible “on top” wherever overlaps happen.

Implement a small API to:

Add a rectangle (with an id, position, size, and a display character).

Render the current board as text.

Click a cell (x, y): if a rectangle exists at that cell, bring the topmost rectangle at that point to the front (top layer).

Move a rectangle to a new position.

Keep the problem interview-style: you can decide and document how to handle edge cases (bounds, invalid moves, clicks on empty cells, etc.).

Example of the Grid

0 0 0 0 0
0 L L 0 0
0 L T T T
0 0 T T T


