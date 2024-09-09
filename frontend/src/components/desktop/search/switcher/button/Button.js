import "./Button.css";
import React from "react";

export function Button(props) {
    const activeButton = (isActive, text) => {
        if (isActive) {
            return <button className={"search-switcher-button-active"}> { text } </button>
        }
        return <button className={"search-switcher-button-inactive"}> { text } </button>
    }

    return ( 
        activeButton(props.isActive, props.text)
    )
}