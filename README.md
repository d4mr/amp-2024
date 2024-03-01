# amp-2024
Backtracking algorithm implemented in TypeScript, transpiled to Golang (using chatgpt lol) to solve Jane Street AMP 2024 Puzzle.

TS version runs for 10 mins+ and crashes because it runs out of memory. Golang solved it in under a second.
I only ran the TS code in Bun, so no idea if it's Bun specific, or if my code is garbage. Would love to know though.
Also how is Go so fast? And why is TS so slow? It's a line by line rewrite pretty much, no clever optimisations involved. The difference is baffling, please open issues with suggestions.

# Puzzle
Original puzzle available here: https://www.janestreet.com/amp24-puzzle/

Place a number between one and six in each square such that no number repeats in the same row or column. Each number represents a New York City skyscraper of the corresponding height. The numbers outside the grid indicate how many skyscrapers are visible in that direction. Note that taller skyscrapers block the view of shorter skyscrapers located behind them.

Once the cityscape is completed, trace the routes horizontally and vertically (but not diagonally) through adjacent skyscrapers to find a hidden message. A route may intersect with itself but not with any other route.

<img width="810" alt="image" src="https://github.com/d4mr/amp-2024/assets/16459486/a8dba8a6-d1b6-4eed-b057-3a2e52750b5f">

Routes:
- 256215
- 26123546
- 425456
- 54643163
- 5632314
