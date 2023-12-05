import React, { useState } from "react";
import { VStack, Spacer, Box, Center } from "@chakra-ui/react";
import UploadFile from "./UploadFile";
import FileDisplay from "./FileDisplay";
import UserInput from "./UserInput";

function Home() {
  // arbitrary value use as a dependency in useEffect in FileDisplay
  const [uploadComplete, setUploadComplete] = useState(0);

  return (
    <Center>
      <Box>
        <VStack spacing={4} align="baseline">
          <UserInput>
            <UploadFile finishUpload={() => setUploadComplete(1)} />
          </UserInput>
          <Spacer />
          <FileDisplay uploadComplete={uploadComplete} />
        </VStack>
      </Box>
    </Center>
  );
}

export default Home;
