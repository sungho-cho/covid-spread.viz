import React, { useEffect, useRef, MutableRefObject } from 'react';

type Callback = (input: void) => void

function useInterval(callback: Callback, delay: number) {
  const savedCallback: MutableRefObject<Callback | undefined> = useRef();

  // Remember the latest callback.
  useEffect(() => {
    savedCallback.current = callback;
  }, [callback]);

  // Set up the interval.
  useEffect(() => {
    function tick() {
      savedCallback.current?.();
    }
    if (delay !== null) {
      let id = setInterval(tick, delay);
      return () => clearInterval(id);
    }
  }, [delay]);
}

export default useInterval;