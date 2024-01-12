import React, { useState } from "react";
import { processTagInput } from "./utils/TagInput";
import { Button } from "./ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "./ui/card";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Spinner } from "@chakra-ui/react";

type getSearchInputType = (input: string[]) => void;

interface Props {
  getSearchInput: getSearchInputType;
}

function TagSearch(props: Props) {
  const [rawTagInput, setTagInput] = useState<{ rawTagInput: string }>({
    rawTagInput: "",
  });
  const [isLoading, setIsLoading] = useState<{ isLoading: Boolean }>({
    isLoading: false,
  });

  function handleSearch(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();

    // clean input and send it to FileDisplay component
    setIsLoading({ isLoading: true });
    const userTags = processTagInput(rawTagInput.rawTagInput);
    props.getSearchInput(userTags);
    setIsLoading({ isLoading: false });
  }

  function handleTagChange(event: React.ChangeEvent<HTMLInputElement>) {
    setTagInput({ rawTagInput: event.target.value });
  }

  return (
    <form onSubmit={handleSearch}>
      <Card>
        <CardHeader>
          <CardTitle>Search</CardTitle>
          <CardDescription>
            Filter your uploaded files by entering tags.
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-2">
          <div className="space-y-1">
            <Label>Tags:</Label>
            <Input
              placeholder="Enter a comma separated list"
              onChange={handleTagChange}
            ></Input>
          </div>
        </CardContent>
        <CardFooter>
          <Button
            type="submit"
            disabled={!!isLoading.isLoading}
            style={{ marginRight: "10px" }}
          >
            Search
          </Button>
          {isLoading.isLoading && <Spinner />}
        </CardFooter>
      </Card>
    </form>
  );
}

export default TagSearch;
