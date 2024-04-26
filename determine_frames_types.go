package determine_frames_types

import (
	"gocv.io/x/gocv"
)

// CalculateSAD вычисляет сумму абсолютных разностей между двумя кадрами
func CalculateSAD(frame1, frame2 gocv.Mat) int {
	diff := gocv.NewMat()
	defer diff.Close()

	gocv.AbsDiff(frame1, frame2, &diff)
	return int(gocv.Sum(diff).Val0)
}

// CalculateOpticalFlowFarneback вычисляет оптический поток между двумя кадрами
func CalculateOpticalFlowFarneback(prevFrame, nextFrame gocv.Mat) gocv.Mat {
	flow := gocv.NewMat()
	gocv.CalcOpticalFlowFarneback(prevFrame, nextFrame, &flow, 0.5, 3, 15, 3, 5, 1.2, 0)
	return flow
}

// CalculateFlowMagnitude вычисляет магнитуду оптического потока
func CalculateFlowMagnitude(flow gocv.Mat) float64 {
	xComp := gocv.NewMat()
	yComp := gocv.NewMat()
	defer xComp.Close()
	defer yComp.Close()

	gocv.Split(flow, &[]gocv.Mat{xComp, yComp})
	magnitude := gocv.NewMat()
	defer magnitude.Close()

	gocv.Magnitude(xComp, yComp, &magnitude)
	return gocv.Mean(magnitude).Val0
}

// GetFrameTypes анализирует список кадров и возвращает список типов кадров
func GetFrameTypes(frames []gocv.Mat, thresholdSAD int, thresholdFlow float64, thresholdBFrame float64) []string {
	if len(frames) < 2 {
		return nil // Необходимо минимум два кадра для анализа
	}

	frameTypes := make([]string, len(frames))
	frameTypes[0] = "I" // Первый кадр считаем I-кадром

	for i := 1; i < len(frames); i++ {
		sad := CalculateSAD(frames[i-1], frames[i])
		flow := CalculateOpticalFlowFarneback(frames[i-1], frames[i])
		flowMagnitude := CalculateFlowMagnitude(flow)

		// Определение типа кадра на основе порогов
		if sad < thresholdSAD && flowMagnitude < thresholdFlow {
			frameTypes[i] = "P"
		} else if flowMagnitude < thresholdBFrame {
			frameTypes[i] = "B"
		} else {
			frameTypes[i] = "I"
		}
	}

	return frameTypes
}
