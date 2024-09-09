import './Order.css';
import { React, useEffect, useState } from 'react';
import { OrderItem } from '../../items/order_item/OrderItem';
import { Header } from '../../header/Header';
import { PopularDestination } from '../../items/popular_destination/PopularDestination';
import { Form } from '../../forms/order/Order';
import { ApiQuery, TokenCheck } from '../../../../services/Api'; 
import { useNavigate, useSearchParams } from 'react-router-dom';

function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
}

function QueryHotel(hotelId, set) {
    ApiQuery('/hotel', {
        method: 'post',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            hotelId: hotelId
        })
    }
    ).then((resp) => {
        if (resp.status != "failed") {
            var hotel = resp.data.hotel;
            set([
                <OrderItem 
                    key='1' 
                    title={hotel.title} 
                    description={hotel.description} 
                    img={hotel.imageUrl} 
                />
            ]);
        }
    });
}

function QueryPopularDestinations(set) {
    ApiQuery('/destinations/popular', {
        method: 'get',
        headers: {
            'Content-Type': 'application/json',
        },
        }
    ).then((resp) => {
        if (resp.status != "failed") {
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
        }
    });
}

export function Order() {
    let navigate = useNavigate();

    var [hotel, setHotel] = useState([]);
    var [popularDestinations, setPopularDestinations] = useState([]);

    const [searchParams, setSearchParams] = useSearchParams();
    var [id, setId] = useState(Number(searchParams.get("hotelId")));
    var [from, setFrom] = useState(searchParams.get("from"));
    var [to, setTo] = useState(searchParams.get("to"));

    useEffect(() => {
        TokenCheck().then((ok) => {
            if (!ok) {
                navigate('/login');
            }
        })

        QueryHotel(id, setHotel);
        QueryPopularDestinations(setPopularDestinations);
    }, []);

    return (
        <div className="order-page">
            <Header />
            <div className="order-wrapper">
                <div className="order-content">
                    <div className="order-content-header">
                        <h1>Бронирование</h1>
                    </div>
                    { hotel }
                    <Form />
                </div>
                <div className="order-popular-destinations-wrapper">
                    <h1 className="order-popular-destinations-title">Популярные направления</h1>
                    <div className="order-popular-destinations">
                        { popularDestinations }
                    </div>
                </div>
            </div>
        </div>
    )
}