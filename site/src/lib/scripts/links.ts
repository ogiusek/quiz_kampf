export class Link {
  text: string;
  redirect: () => void;
  canRender: () => boolean;
}