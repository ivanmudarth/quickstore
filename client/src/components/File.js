import { Box, Text } from "@chakra-ui/react";

function File(props) {
  return (
    <Box key={props.info.key}>
      <img src={props.info.URL} width={200} />
      <Text>{props.info.Name}</Text>
      <Text>{props.info.Size} MB</Text>
      <Text>User Tags: {props.info.UserTags?.join(", ")}</Text>
      <Text>Auto Tags: {props.info.AutoTags?.join(", ")}</Text>
    </Box>
  );
}

export default File;
