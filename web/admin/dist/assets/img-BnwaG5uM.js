function r(t){return t?t.startsWith("data:")||t.startsWith("http")||t.startsWith("/")?t:`data:image/png;base64,${t}`:""}export{r as i};
