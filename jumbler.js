function newName(word) {
  let newWord = word.split("");
  newWord.sort(() => Math.random() - 0.5);
  return newWord.toString().replace(/,/g, "");
}

const name = process.argv[2];
console.log(newName(name));
