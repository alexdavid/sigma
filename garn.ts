import * as garn from "https://garn.io/ts/v0.0.18/mod.ts";

export const sigma = garn.go.mkGoProject({
  description: "Sigma",
  src: ".",
  goVersion: "1.21",
});
