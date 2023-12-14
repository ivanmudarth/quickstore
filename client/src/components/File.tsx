import { Box, Text } from "@chakra-ui/react";

interface Props {
  info: {
    URL: string;
    Name: string;
    Size: string;
    UserTags: string[];
    AutoTags: string[];
  };
}

function File(props: Props) {
  return (
    <Box>
      <img src={props.info.URL} alt={props.info.Name} width={200} />
      <Text>{props.info.Name}</Text>
      <Text>{props.info.Size} MB</Text>
      <Text>User Tags: {props.info.UserTags?.join(", ")}</Text>
      <Text>Auto Tags: {props.info.AutoTags?.join(", ")}</Text>
    </Box>
  );
}

export default File;
