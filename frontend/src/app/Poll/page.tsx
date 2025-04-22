import React from "react";
import Bg from "../../Components/Bg";
import Option from "../../Components/Option";

const Poll: React.FC = () => {
  return (
    <div className="min-h-screen bg-black flex items-start justify-center pt-12">
      <div className="w-full max-w-2xl flex flex-col">
        <h1 className="text-3xl mb-6 mt-32 text-center">Депортация</h1>

        <div className="relative min-h-[320px]">
          {/* варианты ответов */}
          <div className="pr-20">
<Option />
<Option />
<Option />
<Option />

          </div>

          {/* боковые элементы */}
          <div className="absolute top-0 right-0 h-full text-lg flex-col pr-4">
            <div>
              <button className="w-full  h-14 mb-2 items-center justify-center bg-white text-black border rounded-lg
                              transition-colors duration-300 
                              hover:bg-black hover:text-white">
                ↑
              </button>
              <button className="w-full h-14 mb-2 flex items-center justify-center bg-white text-black border rounded-lg
                              transition-colors duration-300
                              hover:bg-black hover:text-white">
                ↓
              </button>
            </div>
            
              <div className=" flex items-center justify-center h-10 text-white mb-6">
                1/20
              </div>
              <button className="w-14 h-14 px-4 py-3 mb-2 flex items-center justify-center bg-white text-black border rounded-lg
                              transition-colors duration-300
                              hover:bg-black hover:text-white">
                Vote
              </button>

          </div>
        </div>
      </div>
      <Bg />
    </div>
  );
};

export default Poll;