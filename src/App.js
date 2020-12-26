import React, { Component } from 'react'
import Axios from 'axios'

import './App.css';
import 'bootstrap/dist/css/bootstrap.css';

import AnimeInfo from './components/AnimeInfo/AnimeInfo';

export default class App extends Component {

  //#region Properties

  state = {
    animeList: [],
    list: [],
    animeId: -1,
    infoShown: false,
    loading: false
  }

  endPoint = 'https://anusic-api.herokuapp.com/api/v1';

  //#endregion

  //#region Lifecycle

  componentDidMount() {
    this.setState({ loading: true });

    Axios.get(`${this.endPoint}/anime`)
      .then(e => {
        this.setState({ animeList: e.data, list: e.data, loading: false });
      })
      .catch(() => this.setState({ loading: false }));
  }

  render() {
    let modal = null;

    if (this.state.infoShown) {
      modal = <AnimeInfo
        endPoint={this.endPoint}
        animeId={this.state.animeId}
        onAnimeClosed={this.onAnimeClosed.bind(this)}
      />;
    }

    return (
      <React.Fragment>
        <nav className="navbar sticky-top navbar-light bg-light px-5">
          <a href="/#" className="navbar-brand">Anusic React</a>
          <div className="form-inline">
            <input
              className="form-control mr-sm-2"
              type="search"
              placeholder="Search"
              aria-label="Search"
              onChange={(e) => {
                if (e.target.value.length > 0) {
                  this.setState({
                    list: this.state.animeList.filter(anime =>
                      anime.name.toLowerCase()
                        .concat((anime.altNames || []).join(' ').toLowerCase())
                        .concat((anime.year || 0).toString())
                        .includes(e.target.value.toLowerCase().trim())
                    )
                  });
                } else {
                  this.setState({ list: this.state.animeList });
                }
              }} />
          </div>
        </nav>

        <main className="p-5">
          <div
            className="alert alert-dark"
            role="alert">
            <b>{this.state.list.length}</b> Anime fetched!

            {
              this.state.loading ?
                <div className="spinner spinner-border float-right" role="status">
                  <span className="visually-hidden"></span>
                </div>
                : ''
            }
          </div>
          <ul
            className="list-group">
            {
              this.state.list.map((e, i) => (
                <a
                  className="list-group-item list-group-item-action"
                  key={i}
                  onClick={() => this.onAnimeClicked(e)}
                >
                  <span className="name">{e.name}</span>
                  {e.year
                    ? <span className="badge badge-secondary float-right">{e.year}</span>
                    : ''}

                  {e.altNames && e.altNames.length > 0
                    ? <p className="alt">{e.altNames.join(",")}</p>
                    : ''}
                </a>
              ))
            }
          </ul>
        </main>

        {modal}
      </React.Fragment>
    );
  }

  //#endregion

  //#region Events

  /**
   * Opens the modal
   * @param {object} anime The anime to open
   */
  onAnimeClicked(anime) {
    this.setState({ infoShown: true, animeId: anime.id });
  }

  /**
   * Closes the modal
   */
  onAnimeClosed() {
    this.setState({ infoShown: false });
  }

  //#endregion
}
