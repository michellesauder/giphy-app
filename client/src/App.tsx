import { useState } from "react";
import "./App.css";
import useSWR from "swr";

export const ENDPOINT = "http://localhost:8080";

interface Gif {
  type: string;
  id: string;
  url: string;
  slug: string;
  bitly_gif_url: string;
  embed_url: string;
}

function App() {
  const [inputValue, setInputValue] = useState("trending");

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  const fetcher = (url: string) =>
    fetch(url).then((response) => response.json());

  const { data, error } = useSWR(`${ENDPOINT}/giphs/${inputValue}`, fetcher);

  if (error) {
    console.error("Error fetching data:", error);
  }

  const giphyData = data?.data || [];
  const embeddedURLs = giphyData.map((gif: Gif) => gif.embed_url);

  return (
    <>
      <form>
        <label>
          Search Field:
          <input type="text" value={inputValue} onChange={handleInputChange} />
        </label>
      </form>
      <div>
        {embeddedURLs?.map((embeddedURL: string, index: number) => (
          <iframe
            key={index} // Make sure to provide a unique key for each iframe
            title={`Giphy Embed ${index}`}
            src={embeddedURL}
            width="480" // Adjust the width and height as needed
            height="360"
            allowFullScreen
          />
        ))}
      </div>
    </>
  );
}

export default App;
