import { ChakraProvider } from "@chakra-ui/react";
import Home from "./components/Home";
/*
App
  Home
    UserInput
      UploadFile - request /upload
      TagSearch - request /search
    FileDisplay - request /display
      File(fileInfo)
*/

function App() {
  return (
    <div>
      <ChakraProvider>
        <Home />
      </ChakraProvider>
    </div>
  );
}

export default App;
