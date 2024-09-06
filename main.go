package main


import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-6
a go version of the gff read that does all the work done by the gff read and some extra.
The input to the go-gffread is a fasta file and either the transcript file or the annotation file in the gff3 format.
The examples of the same are enclosed in the example folder and the sameple files are the same as the gff read.

*/


func main () {

type exonStruct struct {
		name string
		exonAligned string
		start string
		end string
	}

type trancript struct {
		name string
		transcriptAligned string
		start string
		end string
	}

	args := os.Args
	if len(args) != 2 {
		log.Fatal(err.Error())
	  fmt.Println("Arguments list cant be empty")
	}

	genomeRead := os.Args[1:]
	annotationRead := os.Args[2:]
	annotationReadGTF := os.Args[3:]

 fOpen, err := os.Open(genomeRead)
 if err != nil {
		log.Fatal(err.Error())
	}

	fRead := bufio.NewScanner(fOpen)
  for fRead.Scan() {
		line := fRead.Text()
		if string(strings.Split(line, "\t")[2]) == "exon" {
		  exonOpen := []exonStruct{}
			exonOpen = append(exonOpen, exonStruct{
			name: strings.Split(line, "\t")[0],
			exonAligned: strings.Split(line, "\t")[2],
			start : strings.Split(line, "\t")[3],
			end : strings.Split(line, "\t")[4]}
		}
		if string(strings.Split(line, "\t")[2]) == "transcript" {
				transcriptOpen = []transcript{}
			  transcriptOpen = append(transcriptOpen, transcript{
				name : strings.Split(line, "\t")[0],
				transcriptAligned : strings.Split(line, "\t")[2],
				start : strings.Split(line, "\t")[3],
				end : strings.Split(line, "\t")[4]
			})
			}
	}

  type annotateGeneStruct struct {
     name string
		 refseq string
		 orgType string
		 start string
     end string
		 strand string
		 alignedID string
		}

	type annotateexonStruct struct {
		name string
		refseq string
		start string
		end string
		strand string
		alignedID string
	}

	type annotationCDS struct {
		 name string
		 refseq string
		 start string
     end string
		strand string
		alignedID string

	}

	type annotateCDS struct {
		name string
		refseq string
		start string
		end string
		strand string
		alignedID string
	}

  type genomeDetail interface {}

	func (* exonexonStruct{}) exonDraw () {
	 return "interface for drawing the exons"
 }



	f, err := os.Open(annotationRannotationReadGTF)
	if err != nil {
		log.Fatal(err.Error())
	}

	fOpen := bufio.Newscanner(f)
  for fOpen.Scan() {
		line := fOpen.text()
		annotateGene := []annotateGeneStruct{}
		if strings.Split(line, "\t")[2] == "gene" {
			annotateGene = append(annotateGene, annotateGeneStruct{
				name : strings.Split(line, "\t")[0],
				refseq : strings.Split(line, "\t")[2],
			  start : strings.Split(line, "\t")[3],
				end : string.Split(line,"\t")[4],
				strand : strings.Split(line, "\t")[6],
				alignedID : strings.Split(strings.Split(line, "\t")[8], ";")[0]
			})

		if strings.Split(line, "\t")[2] == "exon" {
				annotateExon := []annotateexonStruct{}
				annotateExon = append(annotateExon, annotateexonStruct{
					name : strings.Split(line, "\t")[0],
					refseq : strings.Split(line, "\t")[2],
					start : string.Split(line, "\t")[3],
					end : strings.Split(line, "\t")[4]
					strand : strings.Split(line, "\t")[6],
					aligned : strings.Split(strings.Split(line, "\t")[8], ";")[0]
				})

			}
		if strings.Split(line, "\t")[2] == "CDS" {
				annotatedCDS = []annotateCDS
			}




		}
	}


	}
