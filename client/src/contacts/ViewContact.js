import React, { useState, useEffect } from 'react'
import { API_URL } from '../config';
import { useParams, Link } from "react-router-dom";


export default function ViewContact() {
  let { id } = useParams();
  const [contact, setContact] = useState({})
  const [loading, setLoading] = useState(false)

  

  useEffect(() => {
    const fetchContact = async () => {
      setLoading(true)
      try {
        const response = await fetch(`${API_URL}/contacts/${id}`);
        const json = await response.json();
        setContact(json.data);
        setLoading(false)
      } catch (error) {
        console.log("error", error);
        setLoading(false)
      }
    };

    fetchContact();
  }, [id]);

  return (
    <div>
      {!loading ?
        <div className="flex justify-center">
          <div className="lg:w-1/3 w-full">
            <div className="p-10">
              <div className="mb-10 flex items-center justify-between">
                <Link to="/"><h1 className="font-bold">Go back</h1></Link>
              </div>
              <div className="bg-slate-100 rounded-lg px-5">
                <div className="flex border-b py-4">
                  <div className="mr-4 text-slate-400">Name</div>
                  <div className="text-slate-800 font-medium">{contact.Name}</div>
                </div>
                <div className="flex border-b py-4">
                  <div className="mr-4 text-slate-400">Email</div>
                  <div className="text-slate-800 font-medium">{contact.Email}</div>
                </div>
                <div className="flex border-b py-4">
                  <div className="mr-4 text-slate-400">Contact Number</div>
                  <div className="text-slate-800 font-medium">{contact.MobileNumber}</div>
                </div>
                <div className="flex border-b py-4">
                  <div className="mr-4 text-slate-400">City</div>
                  <div className="text-slate-800 font-medium">{contact.City}</div>
                </div>
                <div className="flex py-4">
                  <div className="mr-4 text-slate-400">Country</div>
                  <div className="text-slate-800 font-medium">{contact.Country}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
        : ''}
    </div>
  )
}