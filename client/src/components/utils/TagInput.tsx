// TODO: move to /components
export function processTagInput(rawInput: string) {
  // for case where user input is "  ,   "
  var commaInput = "," + rawInput + ",";

  var splitInput = commaInput.split(/[ \t]*,[ \t]*/);

  // remove empty string items
  var cleanInput = [];
  for (var tag of splitInput) {
    console.log(splitInput);
    if (tag !== "") {
      cleanInput.push(tag.toLowerCase());
    }
  }

  console.log(cleanInput);
  return cleanInput;
}
