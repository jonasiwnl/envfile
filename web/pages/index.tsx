import Image from 'next/image'

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
      <p>Or scan a document (only works on mobile)</p>
      <input type="file" capture="environment" accept="image/pdf" />
    </main>
  )
}

