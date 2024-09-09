import './Main.css';
import React, { useEffect } from "react";
import { useNavigate } from 'react-router-dom';
import { Header } from "../../header/Header";
import { Search } from "../../search/Search";
import { TokenCheck } from '../../../../services/Api';

export function Main() {
    let navigate = useNavigate();

    useEffect(() => {
        TokenCheck().then((ok) => {
            console.log(ok);
            if (!ok) {
                navigate('/login');
            }
        });
    }, [])

    return (
        <div className="main-page">
            <Header />
            <div className='search-wrapper'>
                <Search />
            </div>
        </div>
    );
}