// By Cburnett - Own work, CC BY-SA 3.0, https://commons.wikimedia.org/w/index.php?curid=1499803

import type { PieceType, Color } from "./chess";

function PieceSVG({
  type,
  color,
}: {
  type: PieceType;
  color: Color;
}) {
  const pieceSVGs = {
    black: {
      king: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <g
            fill="none"
            fillRule="evenodd"
            stroke="#000"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
          >
            <path
              id="path6570"
              fill="none"
              stroke="#000"
              strokeLinejoin="miter"
              d="M22.5 11.63V6"
            ></path>
            <path
              fill="#000"
              strokeLinecap="butt"
              strokeLinejoin="miter"
              d="M22.5 25s4.5-7.5 3-10.5c0 0-1-2.5-3-2.5s-3 2.5-3 2.5c-1.5 3 3 10.5 3 10.5"
            ></path>
            <path
              fill="#000"
              d="M12.5 37c5.5 3.5 14.5 3.5 20 0v-7s9-4.5 6-10.5c-4-6.5-13.5-3.5-16 4V27v-3.5c-2.5-7.5-12-10.5-16-4-3 6 6 10.5 6 10.5z"
            ></path>
            <path strokeLinejoin="miter" d="M20 8h5"></path>
            <path
              stroke="#fff"
              d="M32 29.5s8.5-4 6.03-9.65C34.15 14 25 18 22.5 24.5v2.1-2.1C20 18 10.85 14 6.97 19.85 4.5 25.5 13 29.5 13 29.5"
            ></path>
            <path
              stroke="#fff"
              d="M12.5 30c5.5-3 14.5-3 20 0m-20 3.5c5.5-3 14.5-3 20 0m-20 3.5c5.5-3 14.5-3 20 0"
            ></path>
          </g>
        </svg>
      ),
      queen: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <g
            stroke="#000"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
          >
            <path
              strokeLinecap="butt"
              d="M9 26c8.5-1.5 21-1.5 27 0l2.5-12.5L31 25l-.3-14.1-5.2 13.6-3-14.5-3 14.5-5.2-13.6L14 25 6.5 13.5z"
            ></path>
            <path d="M9 26c0 2 1.5 2 2.5 4 1 1.5 1 1 .5 3.5-1.5 1-1 2.5-1 2.5-1.5 1.5 0 2.5 0 2.5 6.5 1 16.5 1 23 0 0 0 1.5-1 0-2.5 0 0 .5-1.5-1-2.5-.5-2.5-.5-2 .5-3.5 1-2 2.5-2 2.5-4-8.5-1.5-18.5-1.5-27 0"></path>
            <path d="M11.5 30c3.5-1 18.5-1 22 0M12 33.5c6-1 15-1 21 0"></path>
            <circle cx="6" cy="12" r="2"></circle>
            <circle cx="14" cy="9" r="2"></circle>
            <circle cx="22.5" cy="8" r="2"></circle>
            <circle cx="31" cy="9" r="2"></circle>
            <circle cx="39" cy="12" r="2"></circle>
            <path
              fill="none"
              strokeLinecap="butt"
              d="M11 38.5a35 35 1 0 0 23 0"
            ></path>
            <g fill="none" stroke="#fff">
              <path d="M11 29a35 35 1 0 1 23 0M12.5 31.5h20M11.5 34.5a35 35 1 0 0 22 0M10.5 37.5a35 35 1 0 0 24 0"></path>
            </g>
          </g>
        </svg>
      ),
      rook: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <g
            fillRule="evenodd"
            stroke="#000"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
          >
            <path
              strokeLinecap="butt"
              d="M9 39.3h27v-3H9zM12.5 32.3l1.5-2.5h17l1.5 2.5zM12 36.3v-4h21v4z"
            ></path>
            <path
              strokeLinecap="butt"
              strokeLinejoin="miter"
              d="M14 29.8v-13h17v13z"
            ></path>
            <path
              strokeLinecap="butt"
              d="m14 16.8-3-2.5h23l-3 2.5zM11 14.3v-5h4v2h5v-2h5v2h5v-2h4v5z"
            ></path>
            <path
              fill="none"
              stroke="#fff"
              strokeLinejoin="miter"
              strokeWidth="1"
              d="M12 35.8h21M13 31.8h19M14 29.8h17M14 16.8h17M11 14.3h23"
            ></path>
          </g>
        </svg>
      ),
      bishop: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <g
            fill="none"
            fillRule="evenodd"
            stroke="#000"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
          >
            <g fill="#000" strokeLinecap="butt">
              <path d="M9 36.6c3.39-.97 10.11.43 13.5-2 3.39 2.43 10.11 1.03 13.5 2 0 0 1.65.54 3 2-.68.97-1.65.99-3 .5-3.39-.97-10.11.46-13.5-1-3.39 1.46-10.11.03-13.5 1-1.35.49-2.32.47-3-.5 1.35-1.46 3-2 3-2z"></path>
              <path d="M15 32.6c2.5 2.5 12.5 2.5 15 0 .5-1.5 0-2 0-2 0-2.5-2.5-4-2.5-4 5.5-1.5 6-11.5-5-15.5-11 4-10.5 14-5 15.5 0 0-2.5 1.5-2.5 4 0 0-.5.5 0 2z"></path>
              <path d="M25 8.6a2.5 2.5 0 1 1-5 0 2.5 2.5 0 1 1 5 0z"></path>
            </g>
            <path
              stroke="#fff"
              strokeLinejoin="miter"
              d="M17.5 26.6h10m-12.5 4h15m-7.5-14.5v5M20 18.6h5"
            ></path>
          </g>
        </svg>
      ),
      knight: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <g
            fill="none"
            fillRule="evenodd"
            stroke="#000"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
          >
            <path
              fill="#000"
              d="M22 10.3c10.5 1 16.5 8 16 29H15c0-9 10-6.5 8-21"
            ></path>
            <path
              fill="#000"
              d="M24 18.3c.38 2.91-5.55 7.37-8 9-3 2-2.82 4.34-5 4-1.042-.94 1.41-3.04 0-3-1 0 .19 1.23-1 2-1 0-4.003 1-4-4 0-2 6-12 6-12s1.89-1.9 2-3.5c-.73-.994-.5-2-.5-3 1-1 3 2.5 3 2.5h2s.78-1.992 2.5-3c1 0 1 3 1 3"
            ></path>
            <path
              fill="#fff"
              stroke="#fff"
              d="M9.5 25.8a.5.5 0 1 1-1 0 .5.5 0 1 1 1 0"
            ></path>
            <path
              fill="#fff"
              stroke="#fff"
              strokeWidth="1.49997"
              d="M14.933 16.05a.5 1.5 30 1 1-.866-.5.5 1.5 30 1 1 .866.5"
            ></path>
            <path
              fill="#fff"
              stroke="none"
              d="m24.55 10.7-.45 1.45.5.15c3.15 1 5.65 2.49 7.9 6.75s3.25 10.31 2.75 20.25l-.05.5h2.25l.05-.5c.5-10.06-.88-16.85-3.25-21.34s-5.79-6.64-9.19-7.16z"
            ></path>
          </g>
        </svg>
      ),
      pawn: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <path
            stroke="#000"
            strokeLinecap="round"
            strokeWidth="1.5"
            d="M22.5 9c-2.21 0-4 1.79-4 4 0 .89.29 1.71.78 2.38C17.33 16.5 16 18.59 16 21c0 2.03.94 3.84 2.41 5.03-3 1.06-7.41 5.55-7.41 13.47h23c0-7.92-4.41-12.41-7.41-13.47 1.47-1.19 2.41-3 2.41-5.03 0-2.41-1.33-4.5-3.28-5.62.49-.67.78-1.49.78-2.38 0-2.21-1.79-4-4-4z"
          ></path>
        </svg>
      ),
    },
    white: {
      king: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <g
            fill="none"
            fillRule="evenodd"
            stroke="#000"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
          >
            <path
              id="path6570"
              fill="none"
              stroke="#000"
              strokeLinejoin="miter"
              d="M22.5 11.63V6"
            ></path>
            <path
              fill="#000"
              strokeLinecap="butt"
              strokeLinejoin="miter"
              d="M22.5 25s4.5-7.5 3-10.5c0 0-1-2.5-3-2.5s-3 2.5-3 2.5c-1.5 3 3 10.5 3 10.5"
            ></path>
            <path
              fill="#000"
              d="M12.5 37c5.5 3.5 14.5 3.5 20 0v-7s9-4.5 6-10.5c-4-6.5-13.5-3.5-16 4V27v-3.5c-2.5-7.5-12-10.5-16-4-3 6 6 10.5 6 10.5z"
            ></path>
            <path strokeLinejoin="miter" d="M20 8h5"></path>
            <path
              stroke="#fff"
              d="M32 29.5s8.5-4 6.03-9.65C34.15 14 25 18 22.5 24.5v2.1-2.1C20 18 10.85 14 6.97 19.85 4.5 25.5 13 29.5 13 29.5"
            ></path>
            <path
              stroke="#fff"
              d="M12.5 30c5.5-3 14.5-3 20 0m-20 3.5c5.5-3 14.5-3 20 0m-20 3.5c5.5-3 14.5-3 20 0"
            ></path>
          </g>
        </svg>
      ),
      queen: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <g
            stroke="#000"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
          >
            <path
              strokeLinecap="butt"
              d="M9 26c8.5-1.5 21-1.5 27 0l2.5-12.5L31 25l-.3-14.1-5.2 13.6-3-14.5-3 14.5-5.2-13.6L14 25 6.5 13.5z"
            ></path>
            <path d="M9 26c0 2 1.5 2 2.5 4 1 1.5 1 1 .5 3.5-1.5 1-1 2.5-1 2.5-1.5 1.5 0 2.5 0 2.5 6.5 1 16.5 1 23 0 0 0 1.5-1 0-2.5 0 0 .5-1.5-1-2.5-.5-2.5-.5-2 .5-3.5 1-2 2.5-2 2.5-4-8.5-1.5-18.5-1.5-27 0"></path>
            <path d="M11.5 30c3.5-1 18.5-1 22 0M12 33.5c6-1 15-1 21 0"></path>
            <circle cx="6" cy="12" r="2"></circle>
            <circle cx="14" cy="9" r="2"></circle>
            <circle cx="22.5" cy="8" r="2"></circle>
            <circle cx="31" cy="9" r="2"></circle>
            <circle cx="39" cy="12" r="2"></circle>
            <path
              fill="none"
              strokeLinecap="butt"
              d="M11 38.5a35 35 1 0 0 23 0"
            ></path>
            <g fill="none" stroke="#fff">
              <path d="M11 29a35 35 1 0 1 23 0M12.5 31.5h20M11.5 34.5a35 35 1 0 0 22 0M10.5 37.5a35 35 1 0 0 24 0"></path>
            </g>
          </g>
        </svg>
      ),
      rook: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <g
            fillRule="evenodd"
            stroke="#000"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
          >
            <path
              strokeLinecap="butt"
              d="M9 39.3h27v-3H9zM12.5 32.3l1.5-2.5h17l1.5 2.5zM12 36.3v-4h21v4z"
            ></path>
            <path
              strokeLinecap="butt"
              strokeLinejoin="miter"
              d="M14 29.8v-13h17v13z"
            ></path>
            <path
              strokeLinecap="butt"
              d="m14 16.8-3-2.5h23l-3 2.5zM11 14.3v-5h4v2h5v-2h5v2h5v-2h4v5z"
            ></path>
            <path
              fill="none"
              stroke="#fff"
              strokeLinejoin="miter"
              strokeWidth="1"
              d="M12 35.8h21M13 31.8h19M14 29.8h17M14 16.8h17M11 14.3h23"
            ></path>
          </g>
        </svg>
      ),
      bishop: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <g
            fill="none"
            fillRule="evenodd"
            stroke="#000"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
          >
            <g fill="#000" strokeLinecap="butt">
              <path d="M9 36.6c3.39-.97 10.11.43 13.5-2 3.39 2.43 10.11 1.03 13.5 2 0 0 1.65.54 3 2-.68.97-1.65.99-3 .5-3.39-.97-10.11.46-13.5-1-3.39 1.46-10.11.03-13.5 1-1.35.49-2.32.47-3-.5 1.35-1.46 3-2 3-2z"></path>
              <path d="M15 32.6c2.5 2.5 12.5 2.5 15 0 .5-1.5 0-2 0-2 0-2.5-2.5-4-2.5-4 5.5-1.5 6-11.5-5-15.5-11 4-10.5 14-5 15.5 0 0-2.5 1.5-2.5 4 0 0-.5.5 0 2z"></path>
              <path d="M25 8.6a2.5 2.5 0 1 1-5 0 2.5 2.5 0 1 1 5 0z"></path>
            </g>
            <path
              stroke="#fff"
              strokeLinejoin="miter"
              d="M17.5 26.6h10m-12.5 4h15m-7.5-14.5v5M20 18.6h5"
            ></path>
          </g>
        </svg>
      ),
      knight: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <g
            fill="none"
            fillRule="evenodd"
            stroke="#000"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
          >
            <path
              fill="#000"
              d="M22 10.3c10.5 1 16.5 8 16 29H15c0-9 10-6.5 8-21"
            ></path>
            <path
              fill="#000"
              d="M24 18.3c.38 2.91-5.55 7.37-8 9-3 2-2.82 4.34-5 4-1.042-.94 1.41-3.04 0-3-1 0 .19 1.23-1 2-1 0-4.003 1-4-4 0-2 6-12 6-12s1.89-1.9 2-3.5c-.73-.994-.5-2-.5-3 1-1 3 2.5 3 2.5h2s.78-1.992 2.5-3c1 0 1 3 1 3"
            ></path>
            <path
              fill="#fff"
              stroke="#fff"
              d="M9.5 25.8a.5.5 0 1 1-1 0 .5.5 0 1 1 1 0"
            ></path>
            <path
              fill="#fff"
              stroke="#fff"
              strokeWidth="1.49997"
              d="M14.933 16.05a.5 1.5 30 1 1-.866-.5.5 1.5 30 1 1 .866.5"
            ></path>
            <path
              fill="#fff"
              stroke="none"
              d="m24.55 10.7-.45 1.45.5.15c3.15 1 5.65 2.49 7.9 6.75s3.25 10.31 2.75 20.25l-.05.5h2.25l.05-.5c.5-10.06-.88-16.85-3.25-21.34s-5.79-6.64-9.19-7.16z"
            ></path>
          </g>
        </svg>
      ),
      pawn: (
        <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45">
          <path
            stroke="#000"
            strokeLinecap="round"
            strokeWidth="1.5"
            d="M22.5 9c-2.21 0-4 1.79-4 4 0 .89.29 1.71.78 2.38C17.33 16.5 16 18.59 16 21c0 2.03.94 3.84 2.41 5.03-3 1.06-7.41 5.55-7.41 13.47h23c0-7.92-4.41-12.41-7.41-13.47 1.47-1.19 2.41-3 2.41-5.03 0-2.41-1.33-4.5-3.28-5.62.49-.67.78-1.49.78-2.38 0-2.21-1.79-4-4-4z"
          ></path>
        </svg>
      ),
    },
  };

  return pieceSVGs[color][type];
}

export { PieceSVG };
