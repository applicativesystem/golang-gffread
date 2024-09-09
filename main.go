package main


import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
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
    // thinking of implementing the cobra cli here with the flag parse, will implement it.

	genomeRead := os.Args[1:]
	annotationRead := os.Args[2:]
	annotationReadGTF := os.Args[3:]
	annotationmerge := os.Args[4:]
	plotSave := os.Args[4:]

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
			end : strings.Split(line, "\t")[4]
		})
		if string(strings.Split(line, "\t")[2]) == "transcript" {
				transcriptOpen = []transcript{}
        transcriptOpen = append(transcriptOpen,transcript{
				name : strings.Split(line, "\t")[0],
				transcriptAligned : strings.Split(line, "\t")[2],
				start : strings.Split(line, "\t")[3],
				end : strings.Split(line, "\t")[4]
			})
			}
	}

	type transcriptExonAlign struct {
		name string
		exonAligned string
		start string
		end string

	}

	type transcriptAlign struct {
		name string
		transcriptAligned string
		start string
		end string
	}


	for fRead.Scan() {
		line := fRead.Text()
		if string(strings.Split(line, "\t")[2]) == "exon" {
			transcriptExon := []transcriptExonAlign{}
			transcriptExon = append(transcriptExon, exonStruct{
			name: strings.Split(line, "\t")[0],
			exonAligned: strings.Split(line, "\t")[2],
			start : strings.Split(line, "\t")[3],
			end : strings.Split(line, "\t")[4]
		}
		if string(strings.Split(line, "\t")[2]) == "transcript" {
				transcriptAlignOpen = []transcriptALign{}
        transcriptAlignOpen = append(transcriptAlignOpen,transcript{
				name : strings.Split(line, "\t")[0],
				transcriptAligned : strings.Split(line,"\t")[2],
				start : strings.Split(line, "\t")[3],
				end : strings.Split(line, "\t")[4]
			})
			}
	}

	exonAlignEst := []int32{}
	transcriptAlignEst := []int32{}

	for i := range transcriptExon {
	     exonAlignEst = append(exonAlignEst, int(transcriptExon.end)-int(transcriptExon.start) )
	}
	for i := range transcriptALignOpen {
		transcriptAlignEst = append(transcriptAlignEst, int(transcriptAlign.end) - int(transcriptAlignEst.start))
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

	type annotationCDSStruct struct {
		 name string
		 refseq string
		 start string
                 end string
		strand string
		alignedID string

	}

	type annotatetranscriptStruct struct {
		name string
		refseq string
		start string
		end string
		strand string
		alignedID string
	}

	f, err := os.Open(annotationRannotationReadGTF)
	if err != nil {
		log.Fatal(err.Error())
	}

	fOpen := bufio.Newscanner(f)
	for fOpen.Scan() {
		line := fOpen.text()
		if strings.Split(line, "\t")[2] == "gene" {
			annotateGene := []annotateGeneStruct{}
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
					end : strings.Split(line, "\t")[4],
					strand : strings.Split(line, "\t")[6],
					alignedID : strings.Split(strings.Split(line, "\t")[8], ";")[0]
				})
			}
		if strings.Split(line, "\t")[2] == "CDS" {
				annotatedCDS := []annotateCDSStruct{}
				annotateCDS = append(annotateCDS, annotateCDSStruct{
					name : strings.Split(line, "\t")[0],
					refseq : strings.Split(line, "\t")[2],
					start : strings.Split(line, "\t")[3],
					end : strings.Split(line, "\t")[4],
          strand : strings.Split(line, "\t")[6],
					alignedID : strings.Split(strings.Split(line, "\t")[8], ";")[0]
				})
			}
		if strings.Split(line, "\t")[2] == "transcript" {
					annotatetranscript := []annotatetranscriptStruct{}
					annotatetranscript = append(annotatetranscript, annotatetranscriptStruct{
						name : strings.Split(line, "\t")[0],
						refseq : strings.Split(line, "\t")[2],
						start : strings,Split(line, "\t")[3],
						end : strings.Split(line, "\t")[4],
						strand : strings.Split(line, "\t")[6],
						alignedID := strings.Split(strings.Split(line, "\t")[8], ";")[0]
				})
				}
			}
		}

   // extra analysis for gff comparison and length estimates across the gff

	type geneAlign struct {
		name string
		refseq string
		start string
		end string
		strand string
		alignedID string

	}

	type cdsAlign struct {
		name string
		refseq string
		start string
		end string
		strand string
		alignedID string

	}

	type exonAlign struct {
		name string
		refseq string
		start string
		end string
		strand string
		alignedID string

	}

	type transcriptAlign struct {
		name string
		refseq string
		start string
		end string
		strand string
		alignedID string

	}

	// reason why i am declaring it again here is because this will help me in passing this as a complete interface to another package.
	// if you want to remove it then simply make the struct first word capital so that it can be listened as global.
  fOpen := bufio.Newscanner(f)
	for fOpen.Scan() {
		line := fOpen.text()
		if strings.Split(line, "\t")[2] == "gene" {
			geneAlignT := []geneAlign{}
			geneAlignT = append(geneAlign, geneAlign{
				name : strings.Split(line, "\t")[0],
				refseq : strings.Split(line, "\t")[2],
			  start : strings.Split(line, "\t")[3],
				end : string.Split(line,"\t")[4],
				strand : strings.Split(line, "\t")[6],
				alignedID : strings.Split(strings.Split(line,"\t")[8], ";")[0]
			})

		if strings.Split(line, "\t")[2] == "exon" {
				exonAlignT := []exonAlign{}
				exonAlignT = append(exonAlignT,exonAlign{
					name : strings.Split(line, "\t")[0],
					refseq : strings.Split(line, "\t")[2],
					start : string.Split(line, "\t")[3],
					end : strings.Split(line, "\t")[4],
					strand : strings.Split(line, "\t")[6],
					alignedID : strings.Split(strings.Split(line, "\t")[8], ";")[0]
				})
			}
		if strings.Split(line, "\t")[2] == "CDS" {
				cdsAlignT := []cdsAlign{}
				cdsAlignT = append(cdsAlignT, cdsAlign{
					name : strings.Split(line, "\t")[0],
					refseq : strings.Split(line, "\t")[2],
					start : strings.Split(line, "\t")[3],
					end : strings.Split(line, "\t")[4],
          strand : strings.Split(line, "\t")[6],
					alignedID : strings.Split(strings.Split(line, "\t")[8], ";")[0]
				})
			}
		if strings.Split(line, "\t")[2] == "transcript" {
					transcriptAlignT := []transcriptAlign{}
					transcriptAlignT = append(transcriptAlignT, transcriptAlign{
						name : strings.Split(line, "\t")[0],
						refseq : strings.Split(line,"\t")[2],
						start : strings,Split(line,"\t")[3],
						end : strings.Split(line,"\t")[4],
						strand : strings.Split(line,"\t")[6],
						alignedID := strings.Split(strings.Split(line, "\t")[8], ";")[0]
				})
				}
			}
		}
	exonAlignLength := []int32{}
	transcriptAlignLength := []int32{}
	cdsAlignLength := []int32{}
	geneAlignLength := []int32{}

	for i:= range geneAlignT {
		geneAlignedLength = append(geneAlignedLength, int(geneAlign.end)-int(geneAlign.start))
	}
	for i := range cdsAlign {
		cdsAlignLength = append(cdsAlignLength, int(cdsAlign.end) - int(cdsAlign.start))
	}
	for i := range exonAlign {
		exonAlignLength = append(exonAlignLength, int(exonAlign.end) - int(exonAlign.start))
	}
	for i := range transcriptAlign {
		transcriptAlignLength = append(transcriptAlignLength, int(transcriptAlign.end) - int(transcriptAlign.start))
	}

  // implementing the charts 2024--9-7

		bar := charts.NewBar()
		bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title : "Exon plot for the distribution",
			SubTitle: "Plotting the exon distribution"
		}))
		bar.SetAxis(exonAlignT.refseq).AddSeries(exonAligLength)
		f , _ := os.Create("exonplot.html")
		bar.Render(f)


 // merging of several gff aligned across multiple parent genomes

    type annoMergeExon struct {
			name string
			refseq string
			start string
			end string
			strand string
			annomergeE string
		}

		type annoMergeTranscript struct {
			name string
			refseq string
			start string
			end string
			strand string
			annoMergeTran string
		}
		type annoMergeGene struct {
      name string
			refseq string
			start string
			end string
			strand string
			annoMergeGe string
		}
		type annoMergeCDS struct {
			name string
			refseq string
			start string
			end string
			strand string
			annoMergeCD string
		}

		fopen, err := Os.Open(annotationmerge)
		if err != nil {
			log.Fatal(err.Error())
		}
		fread, err := bufio.NewScanner(fopen)
		for fread.Scan() {
			line := fread.Text()
			if strings.HasPrefix(string(line), "@") {
				continue
			} else {
				 if strings.Split(line, "\t")[2] == "exon"{
				annoMergeE := []annoMergeExon{}
				annoMergeG = append(annoMergeT, annoMergeExon{
					name : strings.Split(string(line), "\t")[0],
					refseq : strings.Split(string(line), "\t")[2],
					start : strings.Split(string(line), "\t")[3],
					end : strings.Split(string(line), "\t")[4],
					strand : strings.Split(string(line), "\t")[6],
					annoMergeE : strings.Split(strings.Split(string(line), "\t")[8], ";")[0]
			  })
			 }
			  if strings.Split(line, "\t")[2] == "gene" {
				annoMergeG := []annoMergeGene{}
				annoMergeG = append(annoMergeG, annoMergeGene{
					name : strings.Split(string(line), "\t")[0],
					refseq : strings.Split(string(line), "\t")[2],
					start : strings.Split(string(line), "\t")[3],
					end : strings.Split(string(line), "\t")[4],
					strand : strings.Split(string(line), "\t")[6],
					annoMergeGe : strings.Split(strings.Split(string(line), "\t")[8], ";")[0]
				})
			  }
				if strings.Split(line, "\t")[2] == "transcript" {
			  annoMergeT := []annoMergeTranscript{}
				annoMergeT = append(annoMergeT, annoMergeTranscript{
					name : strings.Split(string(line), "\t")[0],
				  refseq : strings.Split(string(line), "\t")[2],
				  start : strings.Split(string(line), "\t")[3],
				  end : strings.Split(string(line), "\t")[4],
				  strand : = strings.Split(string(line), "\t")[6],
				  annoMergeTran : strings.Split(strings.Split(string(line), "\t")[8], ";")[0]
				})
			  }
				  if strings.Split(line, "\t")[2] == "CDS" {
				 annoMergeC := annoMergeCDS{}
				annoMergeC = append(annoMergeC, annoMergeCDS{
					name : strings.Split(string(line), "\t")[0],
					refseq : strings.Split(string(line), "\t")[2],
					start: strings.Split(string(line), "\t")[3],
					end : strings.Split(string(line), "\t")[4],
					strand : strings.Split(string(line), "\t")[6],
					annoMergCD : strings.Split(strings.Split(string(line), "\t")[8], ";")[0]
				})
			}
		}

			mergedNameGene := []string{}
			mergedstartGene := []string{}
			mergedendGene := []string{}
			megedstrandGene := []string{}
			mergedAlignedGene := []string{}

			for i: range genAlign {
				for j := range annoMergeGene {
					if annoMergeGene.name == geneAlign.name {
						mergedNameGene = append(mergedNameGene, annoMergeGene.name)
						mergedstartGene = append(mergedstartGene, annoMergeGene.start)
						mergedendGene = append(mergedendGene, annoMergGene.end)
						mergedstrandGene = append(mergedstrandGene, annoMergeGene.strand)
						mergedAlignedGene = append(mergedAlignedGene, annoMergeGene.annoMergeGe)
					}
				}

			}

		}


	type genomeDetail interface {}

		func (* exonStruct{}) exonDraw () {
			return "interface for drawing the exons"
	}
}
