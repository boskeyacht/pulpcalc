

export function Card() {
    let handleSimulation = async e => {
        e.preventDefault()
        let tick = e.target.tick.value
        let endTime = e.target.endTime.value
        let frequency = e.target.frequency.value

        console.log(e)
    }

    return (
        <>
            <div className="flex flex-col w-1/3">
                <h1 className="text-center" >Thread Simulation</h1>
                <form className="drop-shadow-xl flex flex-col px-5 pt-7 pb-7 rounded-lg  bg-pink-100" onSubmit={handleSimulation} >
                    <label htmlFor="tick">Tick</label>
                    <input className="shadow appearance-none border rounded w-full py-2 px-3 mb-4 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" type="number" name="tick" placeholder="0" id="tick" required />

                    <label htmlFor="endTime">End Time</label>
                    <input className="shadow appearance-none border rounded w-full py-2 px-3 mb-4 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" type="number" name="endTime" placeholder="0" id="endTime" required />

                    <label htmlFor="frequency">Frequency</label>
                    <input className="shadow appearance-none border rounded w-full py-2 px-3 mb-2 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" type="number" name="frequency" placeholder="0" id="frequency" required />

                    <button className=" bg-pink-500 transition duration-500 hover:bg-pink-400 rounded-2xl mt-5 w-40 py-1 self-center text-white" type="submit">Simulate</button>
                </form>
            </div>
        </>
    )
}