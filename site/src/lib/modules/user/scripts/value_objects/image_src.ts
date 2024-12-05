import img from "../../../../../../public/photos/icons/user.svg";

export function getImageSrc(image: string, secondary_bg: boolean = false): string {
  return image ?
    image :
    secondary_bg ?
      img :
      "https://www.svgrepo.com/show/535711/user.svg"
}