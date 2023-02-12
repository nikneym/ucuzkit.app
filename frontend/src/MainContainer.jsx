export default function MainContainer() {
    let items = [];
    for (let i = 0; i < 100; i++) {
        items[i] =
        <div className="block w-full bg-red-100 mb-4">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla, error hic delectus consequuntur deserunt eveniet perferendis numquam cupiditate quo rerum aliquam, dolor, atque ratione eos voluptatibus officia commodi nemo odit.
        </div>;
    }

    return (
        <>
        <div className="flex justify-center z-10 mt-6 w-full">
            <div className="relative block w-1/2 bg-purple-400 text-black p-4 m-4 rounded-md">
                {items}
            </div>
        </div>
        </>
    );
}
