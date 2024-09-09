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

	// cobra cli will be implemented with flag parse

	args := os.Args
	if len(args) != 2 {
		log.Fatal(err.Error())
	  fmt.Println("Arguments list cant be empty")
	}

	genomeRead := os.Args[1:]
	annotationRead := os.Args[2:]
	annotationReadGTF := os.Args[3:]
	annotationmerge := os.Args[4:]


  // making genome id and seq struct and expecting that your genome is linearized.
	// making a separate struct for the id and the seq so that it can be used separately.

	type genomeIDStruct struct {

		id string
	}

	type genomeSeqStruct struct {
		seq string
	}

	genomeOpen, err := os.Open(genomeRead)
	if err != nil {
		log.Fatal(err.Error())
	}

	genomeScan := bufio.NewScanner(genomeOpen)
	for genomeScan.Scan() {
		line := genomeScan.Text()
		genomeID := []genomeIDStruct{}
		if strings.HasPrefix(string(line), ">") {
			genomeID = append(genomeID, genomeIDStruct{
				id : string(line),
			})
		genomeSeq := []genomeSeqStruct{}
			if !strings.HasPrefix(string(line), ">") {
				genomeSeq = append(genomeSeq, genomeSeqStruct{
					seq : string(line),
				})
			}
		}
	}

  // reading the annotation gff

	fOpen, err := os.Open(annotationRead)
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
			end : strings.Split(line, "\t")[4],
		})
		}
		if string(strings.Split(line, "\t")[2]) == "transcript" {
				transcriptOpen = []transcript{}
        transcriptOpen = append(transcriptOpen,transcript{
				name : strings.Split(line, "\t")[0],
				transcriptAligned : strings.Split(line, "\t")[2],
				start : strings.Split(line, "\t")[3],
				end : strings.Split(line, "\t")[4],
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


	// reading the annotation GTF

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

	// declaring it again here is because this will help me in passing this as a complete interface to another package.

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

  // implementing the charts and they will open as webpages

		bar := charts.NewBar()
		bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title : "Exon plot for the distribution",
			SubTitle: "Plotting the exon distribution"
		}))
		bar.SetAxis(exonAlignT.refseq).AddSeries(exonAlignLength)
		f , _ := os.Create("exonplot.html")
		bar.Render(f)

		bar := charts.NewBar()
		bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title : "Gene plot for the distribution",
			SubTitle: "Plotting the gene distribution"
		}))
		bar.SetAxis(exonAlignT.refseq).AddSeries(geneAlignedLength)
		f , _ := os.Create("geneplot.html")
		bar.Render(f)

		bar := charts.NewBar()
		bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title : "Transcript plot for the distribution",
			SubTitle: "Plotting the transcript distribution"
		}))
		bar.SetAxis(exonAlignT.refseq).AddSeries(transcriptAlignLength)
		f , _ := os.Create("transcriptplot.html")
		bar.Render(f)

		bar := charts.NewBar()
		bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title : "CDS plot for the distribution",
			SubTitle: "Plotting the cds distribution"
		}))
		bar.SetAxis(exonAlignT.refseq).AddSeries(cdsAlignLength)
		f , _ := os.Create("cdsplot.html")
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
			mergedStartGene := []string{}
			mergedEndGene := []string{}
			megedStrandGene := []string{}
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
			mergedNameCDS := []string{}
			mergedStartCDS := []string{}
			mergedEndCDS := []string{}
			mergedStrandCDS := []string{}
			mergedAlignedCDS := []string{}

			for i := range CDSAlign {
				for j := range annoMergeCDS {
					if annoMergeCDS.name == CDSAlign.name {
						mergedNameCDS = append(mergedNameCDS, annoMergeCDS.name)
						mergedStartCDS = append(mergedStartCDS, annoMergeCDS.start)
						mergedEndCDS = append(mergedEndCDS, annoMergeCDS.end)
						mergedStrandCDS = append(mergedStrandCDS, annoMergeCDS.strand)
						mergedAlignedCDS = append(mergedAlignedCDS, annoMergeCDS.annoMergCD)
					}
				}
			}

       mergedNameTranscript := []string{}
			 mergedStartTranscript := []string{}
			 mergedEndTranscript := []string{}
			 mergedStrandTranscript := []string{}
			 mergedAlignedTranscript := []string{}

			 for i := range TranscriptAlign {
					for j := range annoMergeTranscript {
						if annoMergeTranscript.name == TranscriptAlign.name {
							mergedNamename = append(mergedNameTranscript, annoMergeTranscript.name)
							mergedStartTranscript = append(mergedStartTranscript, annoMergeTranscript.start)
						  mergedEndTranscript = append(mergedEndTranscript, annoMergeTranscript.end)
							mergedStrandTranscript = append(mergedStrandTranscript, annoMergeTranscript.strand)
							mergedAlignedTranscript = append(mergedAlignedtranscript, annoMergeTranscript.annoMergeTran)
						}
					}
				}

			mergedNameExon := []string{}
			mergedStartExon := []string{}
			mergedEndExon := []string{}
			mergedStrandExon := []string{}
			mergedAlignedExon := []string{}

     for i := range ExonAlign {
				for j := range annoMergeExon {
					if annoMergeExon.name == ExonAlign.name {
						mergedNameExon = append(mergedNameExon, annoMergeExon.name)
						mergedStartExon = append(mergedNameExon, annoMergeExon.start)
						mergedEndExon = append(mergedNameExon, annoMergeExon.end)
						mergedStrandExon = append(mergedNameExon, annoMergeExon.strand)
						mergedAlignedExon = append(mergedNameExon, annoMergExon.annoMergeE)
					}
				}
			}
}

   // binning of the junctions for the plots.

		 exonBin100 := []int32{}
		 exonBin500 := []int32{}
		 exonBin1000 := []int32{}
		 exonBin1500 := []int32{}
		 exonBin2000 := []int32{}
		 exonBin2500 := []int32{}
		 exonBin3000 := []int32{}
		 exonBin3500 := []int32{}
		 exonBin4000 := []int32{}

		transcriptBin100 := []int32{}
		transcriptBin500 := []int32{}
                transcriptBin1000 := []int32{}
		transcriptBin1500 := []int32{}
		trancriptBin2000 := []int32{}
		transcriptBin2500 := []int32{}
		transcriptBin3000 := []int32{}
		transcriptBin3500 := []int32{}
		trancriptBin40000 := []int32{}

		cdsBin100 := []int32{}
		cdsBin500 := []int32{}
		cdsBin1000 := []int32{}
		cdsBin1500 := []int32{}
		cdsBin2000 := []int32{}
		cdsBin2500 := []int32{}
		cdsBin3000 := []int32{}
		cdsBin3500 := []int32{}
		cdsBin4000 := []int32{}

		geneBin100 := []int32{}
		geneBin500 := []int32{}
		geneBin1000 := []int32{}
		geneBin1500 := []int32{}
		geneBin2000 := []int32{}
		geneBin2500 := []int32{}
		geneBin3000 := []int32{}
		geneBin3500 := []int32{}
		geneBin4000 := []int32{}

		for i := range exonAlignlength {
			if exonAlignnLenth[i] <=100 {
				exonBin100 = append(exonBin100, exonAlignLength[i])
			}
			elseif exonAlignLength[i] >= 100 && exonAlignLength <=500 {
				exonBin500 = append(exonBin500, exonAligLength[i])
			}
			elseif exonAligLength[i] >= 500 && exonAlignLength[i] <=1000 {
				exonBin1000 = append(exonBin1000, exonAlignLength[i])
			}
			elseif exonAlignlength[i] >= 1000 && exonAlignLength[i] <= 1500 {
				exonBin1500 = append(exonBin1500, exonAlignLength[i])
			}
			elseif exonAlignLength[i] >= 1500 && exonAlignLength[i] <= 2000 {
				exonBin2000 = append(exonBin2000, exonAlignLength[i])
			}
			elseif exonAlignLength[i] >= 2000 && exonAlignLength[i] <=2500 {
                                exonBin2500 = append(exonBin2500, exonAlignlength[i])
			}
			elseif exonAlignlength[i] >= 2500 && exonAlignLength[i] <= 3000 {
				exonBin3000 = append(exonBin3000, exonAlignlength[i])
			}
			elseif exonAlignlength[i] >= 3000 && exonAlignlength[i] <= 3500 {
				exonBin35000 = append(exonBin3500, exonAlignLength[i])
			}
			elseif exonAlignLength[i] >= 3500 && exonAlignLength[i] <= 4000 {
				exonBin4000 = append(exonBin4000, exonAlignLength[i])
			}
			elseif {
				continue
			}
 		}

 		for i := range transcriptAlignlength {
			if transcriptAlignnLenth[i] <=100 {
				transcriptBin100 = append(transcriptBin100, transcriptAlignLength[i])
			}
			elseif transcriptAlignLength[i] >= 100 && transcriptAlignLength <=500 {
				transcriptBin500 = append(transcriptBin500, transcriptAligLength[i])
			}
			elseif transcriptAligLength[i] >= 500 && transcriptAlignLength[i] <=1000 {
				transcriptBin1000 = append(transcriptBin1000, transcriptAlignLength[i])
			}
			elseif transcriptAlignlength[i] >= 1000 && transcriptAlignLength[i] <= 1500 {
				transcriptBin1500 = append(transcriptBin1500, transcriptAlignLength[i])
			}
			elseif transcriptAlignLength[i] >= 1500 && transcriptAlignLength[i] <= 2000 {
				transcriptBin2000 = append(transcriptBin2000, transcirptAlignLength[i])
			}
			elseif transcriptAlignLength[i] >= 2000 && transcriptAlignLength[i] <=2500 {
				transcriptBin2500 = append(transcriptBin2500, transcriptAlignlength[i])
			}
			elseif transcriptAlignlength[i] >= 2500 && transcriptAlignLength[i] <= 3000 {
				transcriptBin3000 = append(transcriptBin3000, transcriptAlignlength[i])
			}
			elseif transcriptAlignlength[i] >= 3000 && transcriptAlignlength[i] <= 3500 {
				transcriptBin35000 = append(transcriptBin3500, transcriptAlignLength[i])
			}
		elseif transcriptAlignLength[i] >= 3500 && transcriptAlignLength[i] <= 4000 {
				transcriptBin4000 = append(transcriptBin4000, transcriptAlignLength[i])
			}
			elseif {
				continue
			}
		}

		for i := range geneAlignlength {
			if geneAlignnLenth[i] <=100 {
				geneBin100 = append(geneBin100, geneAligLength[i])
			}
			elseif geneAlignLength[i] >= 100 && geneAlignLength <=500 {
				geneBin500 = append(geneBin500, geneAligLength[i])
			}
			elseif geneAligLength[i] >= 500 && geneAlignLength[i] <=1000 {
				geneBin1000 = append(geneBin1000, geneAlignLength[i])
			}
			elseif geneAlignlength[i] >= 1000 && geneAlignLength[i] <= 1500 {
				geneBin1500 = append(geneBin1500, geneAlignLength[i])
			}
			elseif geneAlignLength[i] >= 1500 && geneAlignLength[i] <= 2000 {
				geneBin2000 = append(geneBin2000, geneAlignLength[i])
			}
			elseif geneAlignLength[i] >= 2000 && geneAlignLength[i] <=2500 {
				geneBin2500 = append(geneBin2500, geneAlignlength[i])
			}
			elseif geneAlignlength[i] >= 2500 && geneAlignLength[i] <= 3000 {
				geneBin3000 = append(geneBin3000, geneAlignlength[i])
			}
			elseif geneAlignlength[i] >= 3000 && geneAlignlength[i] <= 3500 {
				geneBin35000 = append(geneBin3500, geneAlignLength[i])
			}
			elseif geneAlignLength[i] >= 3500 && geneAlignLength[i] <= 4000 {
				geneBin4000 = append(geneBin4000, geneAlignLength[i])
			}
			elseif {
				continue
			}
		}

		for i := range cdsAlignlength {
			if cdsAlignnLenth[i] <=100 {
				cdsBin100 = append(cdsBin100, cdsAligLength[i])
			}
			elseif cdsAlignLength[i] >= 100 && cdsAlignLength <=500 {
				cdsBin500 = append(cdsBin500, cdsAligLength[i])
			}
			elseif cdsAligLength[i] >= 500 && cdsAlignLength[i] <=1000 {
				cdsBin1000 = append(cdsBin1000, cdsAlignLength[i])
			}
			elseif cdsAlignlength[i] >= 1000 && cdsAlignLength[i] <= 1500 {
				cdsBin1500 = append(cdsBin1500, cdsAlignLength[i])
			}
			elseif cdsAlignLength[i] >= 1500 && cdsAlignLength[i] <= 2000 {
				cdsBin2000 = append(cdsBin2000, cdsAlignLength[i])
			}
			elseif cdsAlignLength[i] >= 2000 && cdsAlignLength[i] <=2500 {
				cdsBin2500 = append(cdsBin2500, cdsAlignlength[i])
			}
			elseif cdsAlignlength[i] >= 2500 && cdsAlignLength[i] <= 3000 {
				cdsBin3000 = append(cdsBin3000, cdsAlignlength[i])
			}
			elseif cdsAlignlength[i] >= 3000 && cdsAlignlength[i] <= 3500 {
				cdsBin35000 = append(cdsBin3500, cdsAlignLength[i])
			}
			elseif cdsAlignLength[i] >= 3500 && cdsAlignLength[i] <= 4000 {
				cdsBin4000 = append(cdsBin4000, cdsAlignLength[i])
			}
			elseif {
				continue
			}
		}


   // extract fasta regions from the gtf and the gff

   type genomeIDStruct struct {

	   id string
   }

   type genomeSeqStruct struct {
	   seq string
   }

   genomeOpen, err := os.Open(genomeRead)
   if err != nil {
	   log.Fatal(err.Error())
   }

   genomeScan := bufio.NewScanner(genomeOpen)
   for genomeScan.Scan() {
	   line := genomeScan.Text()
	   genomeID := []genomeIDStruct{}
	   if strings.HasPrefix(string(line), ">") {
		   genomeID = append(genomeID, genomeIDStruct{
			   id : string(line)
		   })
		   genomeSeq := []genomeSeqStruct{}
		   if !strings.HasPrefix(string(line), ">") {
			   genomeSeq = append(genomeSeq, genomeSeqStruct{
				   seq : string(line)
			   })
		   }
	   }
   }


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

   // transcript splice regions extraction from the fasta struct

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

	 geneExtract struct {
		start string
		end string
		seq string
	}

	exonExtract struct {
		start string
		end string
		seq string
	}
	transcriptExtract struct {
		start string
		end string
		seq string
	}
	cdsExtract struct {
		start string
		end string
		seq string
	}

	for i := range geneAlign {
		for j := range genomeSeq {
			if genomeID[i] == geneAlign.name[i] {
		     geneExt = []geneExtract{}
		     geneExt = append(geneExt, geneExtract{
			   start : geneAlign.start[i]
			   end : geneAlign.end[i]
					seq : genomeSeq[i][int(geneAlign.start[i]):int(geneAlign.end[i])]
		})
			}
		}
	}

	for i := range transcriptAlign {
		for j := range genomeSeq {
			if genomeID[i] == transcriptAlign.name[i] {
				transcriptExt = []transcriptExtract{}
				transcriptExt = append(geneExt, geneExtract{
					start : transcriptAlign.start[i]
					end : transcriptAlign.end[i]
					seq : genomeSeq[i][int(transcriptAlign.start[i]):int(transcriptAlign.end[i])]
				})
			}
		}
	}

	for i := range cdsAlign {
		for j := range genomeSeq {
			if genomeID[i] == cdsAlign.name[i] {
				cdsExt = []cdsExtract{}
				cdsExt = append(cdsExt, cdsExtract{
					start : cdsAlign.start[i]
					end : cdsAlign.end[i]
					seq : genomeSeq[i][int(cdsAlign.start[i]):int(cdsAlign.end[i])]
				})
			}
		}
	}

	for i := range exonAlign {
		for j := range genomeSeq {
			if genomeID[i] == exonAlign.name[i] {
				exonExt = []exonExtract{}
				exonExt = append(exonExt, exonExtract{
					start : exonAlign.start[i]
					end : exonAlign.end[i]
					seq : genomeSeq[i][int(exonAlign.start[i]):int(exonAlign.end[i])]
				})
			}
		}
	}
}
}
