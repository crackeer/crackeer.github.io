import React from "react";
import Highlighter from "react-highlight-words";

const makeHighlight = (style) => {
    return ({ children, highlightIndex }) => (
        <strong style={style}>{children}</strong>
    )
}
export default function (props) {
    return <Highlighter
        searchWords={props.searchWords || []}
        autoEscape={true}
        textToHighlight={props.children}
        highlightTag={props.highlightTag || makeHighlight(props.highlightStyle)}
    />
}
