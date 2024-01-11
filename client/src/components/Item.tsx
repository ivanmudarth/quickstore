import Url from "./Url";
// @ts-ignore
import File from "./File";
import { FileInfo, UrlInfo } from "./utils/sharedTypes";

interface Props {
  info: {
    FileInfo: FileInfo;
    UrlInfo: UrlInfo;
    UploadTime: string;
  };
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
