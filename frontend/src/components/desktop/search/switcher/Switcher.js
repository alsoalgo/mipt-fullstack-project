import "./Switcher.css";
import React from "react";
import { Button } from "./button/Button";

export function Switcher({ isHotelPressed, setIsHotelPressed}) {
    const handleHotelClick = (event) => {
        setIsHotelPressed(true);
    } 

    return (
        <div className="search-switcher">
            <div onClick={handleHotelClick}> <Button isActive={isHotelPressed} text="Отель" /> </div>
        </div>  
    );
}