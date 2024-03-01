# amp-2024
Backtracking algorithm implemented in Golang to solve Jane Street AMP 2024 Puzzle

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
