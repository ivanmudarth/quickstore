import React, { useState, useEffect } from "react";
import axios from "axios";
import Item from "./Item";
// @ts-ignore
import Masonry, { ResponsiveMasonry } from "react-responsive-masonry";
import { ItemInfo } from "./utils/sharedTypes";
import { Center, Spinner } from "@chakra-ui/react";
// @ts-ignore
import searchError from "../public/search.png";
// @ts-ignore
import emptyError from "../public/empty.png";

interface Props {
  uploadComplete: number;
  searchInput: string[];
}

// TODO: bug - search a tag. upload a file. display won't update cause searchInput still has a value
// maybe have two different effects with different dependencies
function ItemDisplay(props: Props) {
  const [displayInfo, setDisplay] = useState<{ displayInfo: ItemInfo[] }>({
    displayInfo: [],
  });
  const [isLoading, setIsLoading] = useState<{ isLoading: Boolean }>({
    isLoading: false,
  });
  const containerStyle: React.CSSProperties = {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    paddingTop: "60px",
  };
  const [initiaLoad, setInitialLoad] = useState<{ initiaLoad: Boolean }>({
    initiaLoad: true,
  });

  // is called every time uploadComplete prop is updated
  useEffect(() => {
    handleDisplay(props.searchInput);
  }, [props.uploadComplete, props.searchInput]);

  function handleDisplay(searchInput: string[]) {
    const params = { tags: searchInput };
    const url = "http://localhost:8080/display";
    setIsLoading({ isLoading: true });
    axios
      .get(url, { params })
      .then((response) => {
        console.log(response.data);
        if (response.data == null) {
          setDisplay({ displayInfo: [] });
        } else {
          setDisplay({ displayInfo: response.data });
        }
        setIsLoading({ isLoading: false });
        setInitialLoad({ initiaLoad: false });
      })
      .catch((error) => {
        console.log(error);
        setIsLoading({ isLoading: false });
        setInitialLoad({ initiaLoad: false });
      });
  }

  return (
    <div style={{ paddingLeft: "60px", paddingRight: "60px" }}>
      <h2 className="text-m font-semibold tracking-tight">Uploaded items:</h2>
      <p className="text-sm text-muted-foreground">
        {displayInfo.displayInfo.length} items
      </p>
      <div style={{ paddingTop: "20px", paddingBottom: "20px" }}>
        {initiaLoad.initiaLoad || isLoading.isLoading ? (
          <div style={containerStyle}>
            <Spinner />
          </div>
        ) : displayInfo.displayInfo.length == 0 &&
          props.searchInput.length != 0 ? (
          <div style={containerStyle}>
            <img
              src={searchError}
              alt="Search error"
              style={{ maxHeight: "150px", paddingBottom: "30px" }}
            />
            <h3 className="text-m font-semibold tracking-tight">
              No results found:
            </h3>
            <h2 className="text-m tracking-tight">
              Please try a different query
            </h2>
          </div>
        ) : displayInfo.displayInfo.length == 0 ? (
          <div style={containerStyle}>
            <img
              src={emptyError}
              alt="Empty items"
              style={{ maxHeight: "150px", paddingBottom: "30px" }}
            />
            <h3 className="text-m font-semibold tracking-tight">
              No items uploaded:
            </h3>
            <h2 className="text-m tracking-tight">
              Please upload a file or URL above
            </h2>
          </div>
        ) : (
          <ResponsiveMasonry
            columnsCountBreakPoints={{
              500: 1,
              700: 2,
              950: 3,
              1250: 4,
              1550: 5,
            }}
          >
            <Masonry columnsCount={5} gutter="20px">
              {displayInfo.displayInfo?.map((itemInfo, index) => (
                <Item key={index} info={itemInfo} />
              ))}
            </Masonry>
          </ResponsiveMasonry>
        )}
      </div>
    </div>
  );
}

export default ItemDisplay;
