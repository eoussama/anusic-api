import React, { useEffect, useState } from 'react'
import Axios from 'axios'

import './App.css';
import 'bootstrap/dist/css/bootstrap.css';

function App() {
  const [animeList, setAnimeList] = useState([]);
  const [list, setList] = useState([]);

  useEffect(() => {
    Axios.get('https://anusic-api.herokuapp.com/anime')
      .then(e => {
        setAnimeList(e.data);
        setList(animeList);
      });
  }, [animeList])


  return (
    <React.Fragment>
      <nav className="navbar sticky-top navbar-light bg-light px-5">
        <a className="navbar-brand">Anusic React</a>
        <div className="form-inline">
          <input
            className="form-control mr-sm-2"
            type="search"
            placeholder="Search"
            aria-label="Search"
            onInput={(e) => {
              if (e.target.value.length > 0) {
                console.log(e.target.value);
                setList(animeList.filter(anime =>
                  anime.name.toLowerCase()
                    .concat((anime.altNames || []).join(' ').toLowerCase())
                    .concat((anime.year || 0).toString())
                    .includes(e.target.value.toLowerCase().trim())
                ));
              } else {
                setList(animeList);
              }
            }} />
        </div>
      </nav>

      <main className="p-5">
        <div
          className="alert alert-dark"
          role="alert">
          <b>{list.length}</b> Anime fetched!
        </div>

        <ul
          className="list-group">
          {
            list.map((e, i) => (
              <a
                class="list-group-item list-group-item-action"
                href={`https://myanimelist.net/anime/${e.id}`}
                target="_blank"
                key={i}
              >
                <span class="name">{e.name}</span>

                {e.year
                  ? <span class="badge badge-secondary float-right">{e.year}</span>
                  : ''}

                {e.altNames && e.altNames.length > 0
                  ? <p class="alt">{e.altNames}</p>
                  : ''}
              </a>
            ))
          }
        </ul>
      </main>

    </React.Fragment>
  );
}

export default App;
