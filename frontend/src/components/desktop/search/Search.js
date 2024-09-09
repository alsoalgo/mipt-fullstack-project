import "./Search.css";
import {React, useState, useEffect} from "react";
import { Bar } from "./bar/Bar";
import { Switcher } from "./switcher/Switcher";

export function Search() {
    const [isHotelPressed, setIsHotelPressed] = useState(true);

    useEffect(
        () => {
            setIsHotelPressed(true);
        }, []
    );
    
    return (
        <div className="search">
            <Switcher isHotelPressed={isHotelPressed} setIsHotelPressed={setIsHotelPressed}/>
            <Bar key='search_bar' city='' from='' to='' isHotelPressed={isHotelPressed}/>
        </div>
    )
}