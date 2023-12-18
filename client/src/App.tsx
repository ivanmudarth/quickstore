import { ChakraProvider } from "@chakra-ui/react";
import Home from "./components/Home";
import Header from "./components/Header";
/*
App
  Header
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
        <Header />
        <Home />
      </ChakraProvider>
    </div>
  );
}

export default App;
