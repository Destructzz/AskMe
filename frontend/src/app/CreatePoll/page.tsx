"use client";

import React, { useState } from "react";
import Bg from "../../Components/Bg";

type PollBlock = {
  id: number;
  title: string;
  options: string[];
};

const CreatePoll: React.FC = () => {
  const [polls, setPolls] = useState<PollBlock[]>([
    { id: Date.now(), title: "", options: [""] },
  ]);

  const addPoll = () => {
    setPolls((prev) => [...prev, { id: Date.now(), title: "", options: [""] }]);
  };

  const removePoll = (id: number) => {
    if (polls.length > 1) {
      setPolls((prev) => prev.filter((poll) => poll.id !== id));
    }
  };

  const handleTitleChange = (id: number, value: string) => {
    setPolls((prev) =>
      prev.map((poll) =>
        poll.id === id ? { ...poll, title: value } : poll
      )
    );
  };

  const handleOptionChange = (pollId: number, index: number, value: string) => {
    setPolls((prev) =>
      prev.map((poll) =>
        poll.id === pollId
          ? {
              ...poll,
              options: poll.options.map((opt, i) =>
                i === index ? value : opt
              ),
            }
          : poll
      )
    );
  };

  const addOption = (pollId: number) => {
    setPolls((prev) =>
      prev.map((poll) =>
        poll.id === pollId
          ? { ...poll, options: [...poll.options, ""] }
          : poll
      )
    );
  };

  const removeOption = (pollId: number, index: number) => {
    setPolls((prev) =>
      prev.map((poll) =>
        poll.id === pollId && poll.options.length > 1
          ? {
              ...poll,
              options: poll.options.filter((_, i) => i !== index),
            }
          : poll
      )
    );
  };

  return (
    <div className="min-h-screen bg-black text-white p-6 transition-all overflow-x-hidden">
      <Bg />

      {/* Верхняя панель: "+" слева, "Создать" справа */}
      <div className="flex justify-between items-center mb-6">
        <button
          className="bg-white text-black rounded px-2 py-3 w-12 transition-all hover:scale-110"
          onClick={addPoll}
        >
          +
        </button>
        <button className="border border-white text-white bg-black px-4 py-2 rounded transition-all hover:bg-white hover:text-black">
          Создать
        </button>
      </div>

      {/* Все блоки опросов */}
      <div className="flex flex-col gap-8">
        {polls.map((poll) => (
          <div key={poll.id} className="flex justify-between items-start">
            <div className="flex flex-col gap-4 w-1/2">
              <div className="p-2 rounded">
                {/* Заголовок и минус */}
                <div className="flex items-center gap-2 mb-2">
                  <button
                    className="bg-white text-black rounded px-2 py-3 w-12 transition-all hover:scale-110"
                    onClick={() => removePoll(poll.id)}
                  >
                    -
                  </button>
                  <input
                    placeholder="Заголовок"
                    className="bg-black text-white focus:bg-white focus:text-black focus:border-transparent outline-none rounded w-full px-4 py-2 min-h-[56px] transition-all duration-300"
                    value={poll.title}
                    onChange={(e) => handleTitleChange(poll.id, e.target.value)}
                  />
                </div>

                {/* Варианты */}
                {poll.options.map((option, index) => (
                  <div key={index} className="flex items-center gap-2 mb-2">
                    <input
                      className="px-4 py-2 border border-white outline-none focus:border-transparent rounded w-full bg-black text-white focus:bg-white focus:text-black min-h-[100px] transition-all duration-300"
                      value={option}
                      onChange={(e) =>
                        handleOptionChange(poll.id, index, e.target.value)
                      }
                    />
                    {index === poll.options.length - 1 ? (
                      <div className="flex flex-col gap-1">
                        <button
                          className="bg-white text-black rounded px-2 py-1 w-8 transition-all hover:scale-110"
                          onClick={() => removeOption(poll.id, index)}
                        >
                          -
                        </button>
                        <button
                          className="bg-white text-black rounded px-2 py-1 w-8 transition-all hover:scale-110"
                          onClick={() => addOption(poll.id)}
                        >
                          +
                        </button>
                      </div>
                    ) : (
                      <button
                        className="bg-white text-black rounded px-2 py-1 w-8 transition-all hover:scale-110"
                        onClick={() => removeOption(poll.id, index)}
                      >
                        -
                      </button>
                    )}
                  </div>
                ))}
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default CreatePoll;
