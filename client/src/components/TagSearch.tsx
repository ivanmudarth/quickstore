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

type getSearchInputType = (input: string[]) => void;

interface Props {
  getSearchInput: getSearchInputType;
}

function TagSearch(props: Props) {
  const [rawTagInput, setTagInput] = useState<{ rawTagInput: string }>({
    rawTagInput: "",
  });

  function handleSearch(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();

    // clean input and send it to FileDisplay component
    const userTags = processTagInput(rawTagInput.rawTagInput);
    props.getSearchInput(userTags);
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
          <Button type="submit">Search</Button>
        </CardFooter>
      </Card>
    </form>
  );
}

export default TagSearch;
