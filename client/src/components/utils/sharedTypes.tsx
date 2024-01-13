type finishUploadType = () => void;

export interface FinishUploadProp {
  finishUpload: finishUploadType;
}

export interface FileInfo {
  ID: string;
  S3Key: string;
  FileURL: string;
  Name: string;
  Size: string;
  Type: string;
  UploadTime: string;
  UserTags: string[];
  AutoTags: string[];
}

export interface UrlInfo {
  ID: string;
  ImageURL: string;
  Title: string;
  URL: string;
  UploadTime: string;
  UserTags: string[];
  AutoTags: string[];
}

export interface ItemInfo {
  FileInfo: FileInfo;
  UrlInfo: UrlInfo;
  UploadTime: string;
}
