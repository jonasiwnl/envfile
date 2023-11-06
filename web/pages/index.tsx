export default function Home() {
  return (
    <main
      className="flex min-h-screen flex-col items-center p-24"
    >
      <h1 className="text-2xl">Sink :)</h1>
      
      <br />

      <p>Upload file</p>
      <input type="file" accept="image/pdf" />

      <br />

      <p>Scan a document (only works on mobile)</p>
      <input type="file" capture="environment" accept="image/pdf" />

      <br />

      <p>add a custom filename if you want</p>
      <input type="text" placeholder="custom filename"/>
      <button>Up</button>

      <br />

      <p>Or grab a file with a key</p>
      <input type="text" placeholder="enter key here"/>
      <button>Down</button>
    </main>
  );
}
