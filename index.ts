// top left corner is 0,0
// bottom right is 5,5
// cell is represented by gameBoard[x + (y*6)]

const skylineDefinitions = {
  n: [4, undefined, 3, 3, undefined, undefined],
  s: [undefined, undefined, undefined, 4, undefined, undefined],
  e: [undefined, 4, undefined, undefined, 5, 3],
  w: [3, undefined, undefined, 6, undefined, 3],
};

const gameBoard: (number | undefined)[] = Array(36).fill(undefined);

gameBoard[18] = 1;
gameBoard[19] = 2;
gameBoard[20] = 3;
gameBoard[21] = 4;
gameBoard[22] = 5;
gameBoard[23] = 6;

function solveGameBoard() {
  // fill the first empty spot with the first number
  let nextEmptyIndex = gameBoard.indexOf(undefined);

  if (nextEmptyIndex === -1) {
    return gameBoard;
  }
  console.log(gameBoard);

  for (const i of getValidOptions(nextEmptyIndex)) {
    gameBoard[nextEmptyIndex] = i;
    if (verifyGameBoard()) {
      solveGameBoard();
    }
    gameBoard[nextEmptyIndex] = undefined;
  }
}

function getValidOptions(index: number) {
  let validOptions = [1, 2, 3, 4, 5, 6];

  // if (index === 0) return [3, 4, 5, 6];

  const x = index % 6;
  const y = Math.floor(index / 6);

  for (let i = 0; i < 6; i++) {
    if (gameBoard[x + i * 6]) {
      validOptions = validOptions.filter(
        (option) => option !== gameBoard[x + i * 6]
      );
    }
  }

  for (let i = 0; i < 6; i++) {
    if (gameBoard[i + y * 6]) {
      validOptions = validOptions.filter(
        (option) => option !== gameBoard[i + y * 6]
      );
    }
  }

  return validOptions;
}

function verifyGameBoard() {
  return (
    checkRowsUnique() &&
    checkColumnsUnique() &&
    checkRowsSkylines() &&
    checkColumnsSkylines()
  );
}

function checkRowsUnique() {
  for (let y = 0; y < 6; y++) {
    if (!checkRowUnique(y)) {
      return false;
    }
  }
  return true;
}

function checkColumnsUnique() {
  for (let x = 0; x < 6; x++) {
    if (!checkColumnUnique(x)) {
      return false;
    }
  }
  return true;
}

function checkRowUnique(y: number) {
  const usedNumbers: number[] = [];

  for (let x = 0; x < 6; x++) {
    const cell = gameBoard[x + y * 6];
    if (cell) {
      if (usedNumbers.includes(cell)) {
        return false;
      }
      usedNumbers.push(cell);
    }
  }
  return true;
}

function checkColumnsSkylines() {
  for (let x = 0; x < 6; x++) {
    if (skylineDefinitions.n[x]) {
      let currentMax = 0;
      let visibleCount = 0;

      for (let y = 0; y < 6; y++) {
        const cell = gameBoard[x + y * 6];
        if (cell && cell > currentMax) {
          currentMax = cell;
          visibleCount++;
        }
      }

      if (visibleCount > skylineDefinitions.n[x]!) {
        return false;
      }
    }

    if (skylineDefinitions.s[x]) {
      let currentMax = 0;
      let visibleCount = 0;

      for (let y = 5; y >= 0; y--) {
        const cell = gameBoard[x + y * 6];
        if (cell && cell > currentMax) {
          currentMax = cell;
          visibleCount++;
        }
      }

      if (visibleCount > skylineDefinitions.s[x]!) {
        return false;
      }
    }
  }

  return true;
}

function checkRowsSkylines() {
  for (let y = 0; y < 6; y++) {
    if (skylineDefinitions.e[y]) {
      let currentMax = 0;
      let visibleCount = 0;

      for (let x = 5; x >= 0; x--) {
        const cell = gameBoard[x + y * 6];
        if (cell && cell > currentMax) {
          currentMax = cell;
          visibleCount++;
        }
      }

      if (visibleCount > skylineDefinitions.e[y]!) {
        return false;
      }
    }

    if (skylineDefinitions.w[y]) {
      let currentMax = 0;
      let visibleCount = 0;

      for (let x = 0; x < 6; x++) {
        const cell = gameBoard[x + y * 6];
        if (cell && cell > currentMax) {
          currentMax = cell;
          visibleCount++;
        }
      }

      if (visibleCount > skylineDefinitions.w[y]!) {
        return false;
      }
    }
  }

  return true;
}

function checkColumnUnique(x: number) {
  const usedNumbers: number[] = [];

  for (let y = 0; y < 6; y++) {
    const cell = gameBoard[x + y * 6];
    if (cell) {
      if (usedNumbers.includes(cell)) {
        return false;
      }
      usedNumbers.push(cell);
    }
  }

  return true;
}

console.log(solveGameBoard());
