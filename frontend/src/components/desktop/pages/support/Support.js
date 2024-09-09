import './Support.css';
import React, { useState, useEffect } from "react";
import { Header } from "../../header/Header";
import { Form } from '../../forms/support/Support';
import { PopularDestination } from '../../items/popular_destination/PopularDestination';
import { ApiQuery, TokenCheck } from '../../../../services/Api';
import { useNavigate } from 'react-router-dom';

function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
}

function QueryPopularDestinations(set) {
    ApiQuery('/destinations/popular', {
        method: 'get',
        headers: {
            'Content-Type': 'application/json',
        },
        }
    ).then((resp) => {
        var destinations = resp.data.destinations;
        set(destinations.map(
            (item, index) =>
            <PopularDestination 
                key={index + 1} 
                title={capitalizeFirstLetter(item.city)} 
                description={'от ' + item.cost + 'Р'} 
                img={item.imageUrl} 
            />
        ));
    });
}

export function Support() {
    let navigate = useNavigate();
    var [popularDestinations, setPopularDestinations] = useState([]);

    useEffect(
        () => {
            TokenCheck().then((ok) => {
                if (!ok) {
                    navigate('/login');
                }
            });

            QueryPopularDestinations(setPopularDestinations);
        }, []);

    return (
        <div className="support-page"> 
            <Header />
            <div className="support-wrapper">
                <div className="support-form">
                    <Form />
                </div>
                <div className="support-popular-destinations-wrapper">
                    <h1 className="support-popular-destinations-title">Популярные направления</h1>
                    <div className="support-popular-destinations">
                        { popularDestinations }
                    </div>
                </div>
            </div>
        </div>
    );
}