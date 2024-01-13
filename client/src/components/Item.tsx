import Url from "./Url";
// @ts-ignore
import File from "./File";
import { ItemInfo } from "./utils/sharedTypes";

interface Props {
  info: ItemInfo;
}

function Item(props: Props) {
  return (
    <>
      {props.info.FileInfo == null ? (
        <Url urlinfo={props.info.UrlInfo} />
      ) : (
        <File fileinfo={props.info.FileInfo} />
      )}
    </>
  );
}

export default Item;
