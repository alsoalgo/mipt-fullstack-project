import './Hotels.css';
import React from 'react';
import { useState, useEffect } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { Header } from '../../header/Header';
import { Bar } from '../../search/bar/Bar';
import { PopularDestination } from '../../items/popular_destination/PopularDestination';
import { SearchResult } from './search/result/SearchResult';
import { Switcher } from '../../search/switcher/Switcher';
import { ApiQuery, TokenCheck } from '../../../../services/Api';

function QuerySearch(city, from, to, set) {
    ApiQuery('/search?' + 'city=' + city + '&from=' + from + '&to=' + to, {
        method: 'post',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({})
    }
    ).then((resp) => {
        if (resp.status != "failed") {
            var hotels = resp.data.hotels;
            set(hotels.map(
                (item, index) =>
                <SearchResult 
                    key={index + 1} 
                    hotelId={item.id}
                    from={from}
                    to={to}
                    title={item.title} 
                    description={item.description} 
                    img={item.imageUrl} 
                    />
            ));
        }
    });
}

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

function validateSearchParams(city, from, to) {
    return (!city || !from || !to);
}

export function Hotels() {
    let navigate = useNavigate();
    var [searchResults, setSearchResults] = useState([]);
    var [popularDestinations, setPopularDestinations] = useState([]);
    const [searchParams, setSearchParams] = useSearchParams();
    var [city, setCity] = useState(searchParams.get("city"));
    var [from, setFrom] = useState(searchParams.get("from"));
    var [to, setTo] = useState(searchParams.get("to"));
    
    useEffect(
        () => {
            TokenCheck().then((ok) => {
                if (!ok) {
                    navigate('/login');
                }
            });

            QuerySearch(
                searchParams.get("city"), 
                searchParams.get("from"), 
                searchParams.get("to"),
                setSearchResults
            );
            QueryPopularDestinations(setPopularDestinations);
        }, [window.location.search]
    );

    const [isHotelPressed, setIsHotelPressed] = useState(true);

    return (
        <div className='hotels-page'>
            <Header />
            <div className='hotels-search'>
                <Switcher isHotelPressed={isHotelPressed} setIsHotelPressed={setIsHotelPressed}/>
                <Bar key='hotels_bar' city={city} from={from} to={to} isHotelPressed={isHotelPressed}/>
            </div>
            <div className='hotels-wrapper'>
                <div key={city + from + to} className='hotels-search-results'>
                    { searchResults }
                </div>
                <div className='hotels-popular-destinations-wrapper'> 
                    <h1 className='hotels-popular-destinations-title'>Популярные направления</h1>
                    <div className='hotels-popular-destinations'>
                        { popularDestinations }
                    </div>
                </div>
            </div>
        </div>
    );
}
