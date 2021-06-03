import * as React from 'react';

function SvgReferrer(props: React.SVGProps<SVGSVGElement>) {
    return (
        <svg
            width="1em"
            height="1em"
            viewBox="0 0 64 50"
            fill="current"
            xmlns="http://www.w3.org/2000/svg"
            {...props}
        >
            <path d="M.3 37.865c1.1-5.9 3.2-9.6 7.2-13.9 5.6-5.6 13.4-9.4 21.7-11v-7.5c0-1.6.5-3.2 1.6-4 2.1-1.9 5.1-1.9 7-.3l24.1 19.3c1.3 1.1 2.1 2.7 2.1 4.3 0 1.6-.8 3.2-1.9 4.3L38 48.565c-1.1.8-2.1 1.1-3.2 1.1-.8 0-1.6-.3-2.4-.5-1.9-.8-2.9-2.7-2.9-4.8v-6.2c-7.8 1.3-15.5 5.4-21.2 8.8-1.6 1.1-3.7 1.1-5.6 0-1.6-1.1-2.7-2.9-2.7-4.8 0-2.1 0-3.5.3-4.3zm31.3-5.6l2.9-.3v12.1l24.1-19.3-24.1-19.3v12.4l-2.4.3c-8 .8-15.5 4.3-20.9 9.6-3.2 3.5-4.8 6.4-5.9 11.2 0 .3-.3 1.3-.3 3.2 6.8-4.3 17-9.3 26.6-9.9z" />
        </svg>
    );
}

export default SvgReferrer;
