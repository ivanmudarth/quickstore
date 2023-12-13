// TODO: move to /components
export function processTagInput(rawInput) {
  // for case where user input is "  ,   "
  var rawInput = "," + rawInput + ",";

  var splitInput = rawInput.split(/[ \t]*,[ \t]*/);

  // remove empty string items
  var cleanInput = [];
  for (var tag of splitInput) {
    console.log(splitInput);
    if (tag != "") {
      cleanInput.push(tag.toLowerCase());
    }
  }

  console.log(cleanInput);
  return cleanInput;
}
