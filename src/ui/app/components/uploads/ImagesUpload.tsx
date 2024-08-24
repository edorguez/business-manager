"use client";

import { Icon } from "@iconify/react";
import Image from "next/image";
import { ChangeEvent, useState } from "react";
import { useToast } from "@chakra-ui/react";
import SimpleToast from "../toasts/SimpleToast";

interface ImagesUploadProps {
  showAddImage: boolean;
  maxImagesNumber?: number;
  maxImageSizeMb?: number;
}

const ImagesUpload: React.FC<ImagesUploadProps> = ({
  showAddImage,
  maxImagesNumber = 5,
  maxImageSizeMb = 2.2, // Some files of 2mb are rounded to 2.2
}) => {
  const toast = useToast();
  const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);
  const [imagePreviewUrls, setImagePreviewUrls] = useState<string[]>([]);

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files; // Get the files selected by the user

    if (files) {
      if (!areFilesValid(files)) return;

      // Get files to upload that are not repeated
      let filesToLoad: any[] = Array.from(files).filter(
        (x) => !uploadedFiles.map((y) => y.name).includes(x.name)
      );

      if (filesToLoad.length > 0) {
        // Set the uploaded files
        setUploadedFiles((prevFiles) => [...prevFiles, ...filesToLoad]);
        // Read and set preview URLs for each file
        filesToLoad.forEach((file: any) => {
          const reader = new FileReader();
          reader.readAsDataURL(file);
          reader.onloadend = () => {
            // Once the file is read, create a Blob URL for the image preview
            if (reader.result) {
              const blobUrl = URL.createObjectURL(file);
              setImagePreviewUrls((prevUrls) => [...prevUrls, blobUrl]);
            }
          };
        });
      }
    }
  };

  const areFilesValid = (files: FileList): boolean => {
    const arrayFiles: File[] = Array.from(files);

    if (uploadedFiles.length + files.length - 1 >= maxImagesNumber) {
      toast({
        position: "top-right",
        duration: 6000,
        render: () => (
          <SimpleToast
            title="Error al subir imagen"
            description={`Solo puedes subir un máximo de ${maxImagesNumber} imágenes.`}
            status="error"
          />
        ),
      });
      return false;
    }

    const isIncorrectFileImage: boolean = arrayFiles.some(
      (x: File) => !x.name.match(/\.(jpg|jpeg|png)$/)
    );
    if (isIncorrectFileImage) {
      toast({
        position: "top-right",
        duration: 6000,
        render: () => (
          <SimpleToast
            title="Error al subir imagen"
            description="Solo puedes subir imágenes (PNG, JPG, JPEG)"
            status="error"
          />
        ),
      });
      return false;
    }

    const isFileSizeBig: boolean = arrayFiles.some(
      (x: File) => x.size / 1000000 > maxImageSizeMb
    ); // File size is in kilobytes and we compare with megabytes
    if (isFileSizeBig) {
      toast({
        position: "top-right",
        duration: 6000,
        render: () => (
          <SimpleToast
            title="Error al subir imagen"
            description={`El tamaño de una imagen es mayor al permitido de ${Math.round(
              maxImageSizeMb
            )}mb.`}
            status="error"
          />
        ),
      });
      return false;
    }

    return true;
  };

  const handleRemoveImage = (index: number) => {
    const updatedFiles = [...uploadedFiles];
    const updatedPreviewUrls = [...imagePreviewUrls];

    updatedFiles.splice(index, 1); // Remove the file from the array
    updatedPreviewUrls.splice(index, 1); // Remove the preview URL from the array

    setUploadedFiles(updatedFiles);
    setImagePreviewUrls(updatedPreviewUrls);
  };

  return (
    <>
      {showAddImage && (
        <div>
          <div className="flex justify-center">
            <label
              htmlFor="files"
              className="
            text-sm
            bg-maincolorhov
            text-maincolor
            hover:bg-thirdcolorhov
            hover:text-thirdcolor
            transition
            duration-150
            cursor-pointer
            rounded-full
            font-semibold
            py-2
            px-4
            border-0
            flex
            items-center
            select-none
          "
            >
              <Icon icon="icon-park-outline:upload-picture" className="mr-2" />
              Subir Imagen
            </label>
            <input
              id="files"
              multiple
              onChange={handleFileChange}
              className="hidden"
              type="file"
              accept="image/png, image/gif, image/*"
            />
          </div>
          <div className="w-100 flex justify-center mt-2">
            <small className="text-center text-slate-500">
              Máximo {maxImagesNumber} imágenes (PNG, JPG, JPEG) y{" "}
              {Math.round(maxImageSizeMb)}mb de tamaño
            </small>
          </div>

          <hr className="my-5" />
        </div>
      )}

      {uploadedFiles.length > 0 && (
        <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4 place-items-center">
          {uploadedFiles.map((file, index) => (
            <div key={index} className="relative w-[100px]">
              <button
                onClick={() => handleRemoveImage(index)}
                className="absolute top-[-10px] right-[-10px] rounded-full bg-thirdcolor hover:bg-rose-600 text-white text-lg p-1 transition duration-100"
              >
                <Icon icon="material-symbols:close" />
              </button>
              {imagePreviewUrls[index] && (
                <Image
                  src={imagePreviewUrls[index]}
                  alt=""
                  width={100}
                  height={100}
                />
              )}
            </div>
          ))}
        </div>
      )}

      {uploadedFiles.length === 0 && (
        <div className="text-center">
          <h1 className="font-bold text-md text-thirdcolor">
            Ninguna imagen subida
          </h1>
          <br />
          <span className="text-sm">
            No has subido ninguna imagen, para hacerlo presiona el botón
          </span>
          <br />
          <span className="text-sm">
            <b>&ldquo;Subir Imagen&rdquo;</b>
          </span>
        </div>
      )}
    </>
  );
};

export default ImagesUpload;
