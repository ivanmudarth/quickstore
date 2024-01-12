import React, { useState, useEffect } from "react";
import axios from "axios";
import Item from "./Item";
// @ts-ignore
import Masonry, { ResponsiveMasonry } from "react-responsive-masonry";

interface Props {
  uploadComplete: number;
  searchInput: string[];
}

// TODO: bug - search a tag. upload a file. display won't update cause searchInput still has a value
// maybe have two different effects with different dependencies
// TODO: display message if no search results or no files uploaded yet
function ItemDisplay(props: Props) {
  const [displayInfo, setDisplay] = useState([]); // TODO: add type def

  // is called every time uploadComplete prop is updated
  useEffect(() => {
    handleDisplay(props.searchInput);
  }, [props.uploadComplete, props.searchInput]);

  function handleDisplay(searchInput: string[]) {
    const params = { tags: searchInput };
    const url = "http://localhost:8080/display";
    axios
      .get(url, { params })
      .then((response) => {
        console.log(response.data);
        setDisplay(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }

  return (
    <div style={{ paddingLeft: "60px", paddingRight: "60px" }}>
      <h2 className="text-m font-semibold tracking-tight">Uploaded items:</h2>
      <p className="text-sm text-muted-foreground">
        {displayInfo.length} items
      </p>
      <div style={{ paddingTop: "20px", paddingBottom: "20px" }}>
        <ResponsiveMasonry
          columnsCountBreakPoints={{ 500: 1, 700: 2, 950: 3, 1250: 4, 1550: 5 }}
        >
          <Masonry columnsCount={5} gutter="20px">
            {displayInfo?.map((itemInfo, index) => (
              <Item key={index} info={itemInfo} />
            ))}
          </Masonry>
        </ResponsiveMasonry>
      </div>
    </div>
  );
}

export default ItemDisplay;