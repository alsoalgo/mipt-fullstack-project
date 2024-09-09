import './SearchResult.css';
import { React } from 'react';
import { useNavigate } from 'react-router-dom';

export function SearchResult({hotelId, from, to, title, description, img}) {
    let navigate = useNavigate();

    const book = (event) => {
        event.preventDefault();

        navigate(
            '/order?hotelId=' + hotelId + '&from=' + from + '&to=' + to
        );
    }

    return (
        <div className='hotels-search-result'>
            <img className='hotels-search-result-bg' alt='bg' src={img}></img>
            <div className='hotels-search-result-description-wrapper'>
                <h1 className='hotels-search-result-description-title'>{title}</h1>
                <p className='hotels-search-result-description-sub'>{description}</p>
                <div className='hotels-search-result-description-book-wrapper'> 
                    <button className='hotels-search-result-description-book' onClick={book}>Забронировать</button>
                </div>
            </div>
        </div>
    )
}