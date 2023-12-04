import { VStack, Spacer, Box, Center } from "@chakra-ui/react";
import UploadFile from "./UploadFile";
import FileDisplay from "./FileDisplay";

function Home() {
  return (
    <Center>
      <Box>
        <VStack spacing={4} align="baseline">
          <UploadFile />
          <Spacer />
          <FileDisplay />
        </VStack>
      </Box>
    </Center>
  );
}

export default Home;
